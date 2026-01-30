package learn_transaction

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/IBM/sarama"
)

// var enableTransaction = false
var enableTransaction = true

func TestSaramaTransaction1(t *testing.T) {
	// 配置
	kc := sarama.NewConfig()

	// 消费隔离级别，默认是ReadUncommitted，设置为ReadCommitted可以读取已提交的消息
	kc.Consumer.IsolationLevel = sarama.ReadCommitted
	kc.Consumer.Offsets.Initial = sarama.OffsetOldest // 从最早偏移量开始消费

	kc.Producer.Transaction.ID = "hzw-transaction-1" // 事务ID，用于标识事务
	kc.Producer.Idempotent = true                    // 开启幂等性，确保事务消息的重试不被重复处理
	kc.Producer.RequiredAcks = sarama.WaitForAll     // 等待所有副本确认
	kc.Producer.Return.Successes = true              // 开启返回成功消息，确保事务消息的重试不被重复处理
	kc.Net.MaxOpenRequests = 1                       // 最大并发请求数，默认是5，必须设置为1可以确保事务消息的顺序性

	err := kc.Validate()
	if err != nil {
		t.Fatalf("Kafka配置校验失败: %v", err)
	}

	// client
	client, err := sarama.NewClient(brokers, kc)
	if err != nil {
		t.Fatalf("创建Kafka客户端失败: %v", err)
	}
	defer client.Close()

	// consumer group
	cg, err := sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		t.Fatalf("创建Kafka消费者组失败: %v", err)
	}
	defer cg.Close()

	// producer
	// producer, err := sarama.NewSyncProducerFromClient(client)
	producer, err := sarama.NewSyncProducer(brokers, kc)
	if err != nil {
		t.Fatalf("创建Kafka生产者失败: %v", err)
	}

	// kafkaConfig := sarama.NewConfig()
	// kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	// kafkaConfig.Producer.Retry.Max = 3
	// kafkaConfig.Producer.Return.Successes = true
	// kafkaConfig.Producer.Return.Errors = true
	// producer, err := sarama.NewSyncProducer(brokers, kafkaConfig)
	if err != nil {
		t.Fatalf("创建Kafka生产者失败: %v", err)
	}
	defer producer.Close()

	resultChan := make(chan *sarama.ProducerMessage, 3)
	// go func() {
	// 	for msg := range resultChan {
	// 		start := time.Now()
	// 		partition, offset, err := producer.SendMessage(msg)
	// 		if err != nil {
	// 			slog.Error("发送消息到结果主题失败:", "err", err)
	// 		}
	// 		slog.Info(fmt.Sprintf("发送结果消息, partition: %d, offset: %d, 耗时: %v\n", partition, offset, time.Since(start)))
	// 	}
	// }()

	for {
		cg.Consume(
			context.Background(),
			[]string{topic},
			&_st1_ConsumerGroupHandler{
				client:     client,
				producer:   producer,
				producers:  make(map[int32]sarama.SyncProducer),
				resultChan: resultChan,
			},
		)
	}
}

type _st1_ConsumerGroupHandler struct {
	client     sarama.Client
	producer   sarama.SyncProducer
	producers  map[int32]sarama.SyncProducer
	resultChan chan *sarama.ProducerMessage
}

func (h *_st1_ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	fmt.Printf("Setup\n")

	return nil
}

func (h *_st1_ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *_st1_ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Printf("开始监听:%s[%d] enableTransaction:%v\n", claim.Topic(), claim.Partition(), enableTransaction)

	// producer, ok := h.producers[claim.Partition()]
	// if !ok {
	// 	var err error
	// 	slog.Info(fmt.Sprintf("创建分区生产者:%s[%d]\n", claim.Topic(), claim.Partition()))
	// 	producer, err = sarama.NewSyncProducerFromClient(h.client)
	// 	if err != nil {
	// 		slog.Error(fmt.Sprintf("创建分区生产者:%s[%d]失败: %v\n", claim.Topic(), claim.Partition(), err))
	// 		return err
	// 	}
	// 	h.producers[claim.Partition()] = producer
	// }
	producer := h.producer

	for msg := range claim.Messages() {
		// 耗时切片
		var consumeDurations []time.Duration
		var err error
		start := time.Now()

		if enableTransaction {
			// 开始事务
			err = producer.BeginTxn()
			if err != nil {
				slog.Error("开始事务失败:", "err", err)
				return err
			}
		}
		consumeDurations = append(consumeDurations, time.Since(start))
		start = time.Now()

		// 处理消息
		slog.Info(fmt.Sprintf("收到消息: %s\n", string(msg.Value)))

		restMsg := &sarama.ProducerMessage{
			Topic: resultTopic,
			Key:   sarama.ByteEncoder(msg.Key),
			Value: sarama.StringEncoder("处理成功"),
		}
		consumeDurations = append(consumeDurations, time.Since(start))
		start = time.Now()

		// 发送结果消息
		// h.resultChan <- restMsg
		_, _, err = producer.SendMessage(restMsg) // TODO 500ms
		if err != nil {
			slog.Error("发送结果消息失败:", "err", err)
			if enableTransaction {
				// 消息失败，回滚事务
				producer.AbortTxn()
			}
			return err
		}

		consumeDurations = append(consumeDurations, time.Since(start))
		start = time.Now()

		// 将偏移量提交纳入事务
		if !enableTransaction {
			session.MarkMessage(msg, "") // 不会受BeginTxn影响，会最终提交偏移量
		} else {
			err = producer.AddMessageToTxn(msg, groupID, nil)
			if err != nil {
				slog.Error("添加消息到事务失败:", "err", err)
				if enableTransaction {
					// 消息失败，回滚事务
					producer.AbortTxn()
				}
				return err
			}
		}
		consumeDurations = append(consumeDurations, time.Since(start))
		start = time.Now()

		if enableTransaction {
			// 提交事务（包含消息发送和偏移量提交）
			err = producer.CommitTxn() // TODO 耗时1s
			if err != nil {
				slog.Error("提交事务失败:", "err", err)
				// 消息失败，回滚事务
				producer.AbortTxn() // 重要：此会导致消息会被重复处理，且会发生消费者组重平衡 ！！！
				return err
			}
		}
		consumeDurations = append(consumeDurations, time.Since(start))
		start = time.Now()
		// 打印耗时
		slog.Info(fmt.Sprintf("处理消息耗时: %v\n", consumeDurations))
	}
	return nil
}
