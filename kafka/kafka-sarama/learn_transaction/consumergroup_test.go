package learn_transaction

import (
	"context"
	"fmt"
	"log/slog"
	"testing"

	"github.com/IBM/sarama"
)

func TestSaramaConsumerGroup1(t *testing.T) {
	// 配置
	kc := sarama.NewConfig()

	// 消费隔离级别，默认是ReadUncommitted，设置为ReadCommitted可以读取已提交的消息
	kc.Consumer.IsolationLevel = sarama.ReadCommitted
	kc.Consumer.Offsets.Initial = sarama.OffsetOldest // 从最早偏移量开始消费

	// kc.Producer.Transaction.ID = "hzw-transaction-2" // 事务ID，用于标识事务
	// kc.Producer.Idempotent = true // 开启幂等性，确保事务消息的重试不被重复处理
	// kc.Producer.RequiredAcks = sarama.WaitForAll // 等待所有副本确认
	// kc.Producer.RequiredAcks = sarama.WaitForAll // 等待所有副本确认
	kc.Producer.Return.Successes = true // 开启返回成功消息，确保事务消息的重试不被重复处理
	// kc.Net.MaxOpenRequests = 1                   // 最大并发请求数，默认是5，必须设置为1可以确保事务消息的顺序性
	kc.Net.MaxOpenRequests = 5

	// client
	client, err := sarama.NewClient(brokers, kc)
	if err != nil {
		t.Fatalf("创建Kafka客户端失败: %v", err)
	}
	defer client.Close()

	// consumer group
	c, err := sarama.NewConsumerGroupFromClient("hellohzw", client)
	if err != nil {
		t.Fatalf("创建Kafka消费者失败: %v", err)
	}
	defer c.Close()

	// producer
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		t.Fatalf("创建Kafka生产者失败: %v", err)
	}
	defer producer.Close()
	for {
		c.Consume(
			context.Background(),
			[]string{topic},
			&_scg1_ConsumerGroupHandler{producer: producer},
		)
	}
}

type _scg1_ConsumerGroupHandler struct {
	producer sarama.SyncProducer
}

func (h *_scg1_ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *_scg1_ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *_scg1_ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Printf("开始监听:%s[%d]\n", claim.Topic(), claim.Partition())
	for msg := range claim.Messages() {

		// 处理消息
		slog.Info(fmt.Sprintf("收到消息: %s\n", string(msg.Value)))

		restMsg := &sarama.ProducerMessage{
			Topic: resultTopic,
			Key:   sarama.ByteEncoder(msg.Key),
			Value: sarama.StringEncoder("处理成功"),
		}

		// 发送结果消息
		_, _, err := h.producer.SendMessage(restMsg)
		if err != nil {
			slog.Error(fmt.Sprintf("发送结果消息失败: %v", err))
			return err
		}

		// 将偏移量提交纳入事务
		session.MarkMessage(msg, "") // 不会受BeginTxn影响，会最终提交偏移量
		// err = h.producer.AddMessageToTxn(msg, groupID, nil)
		// if err != nil {
		// 	return err
		// }

	}
	return nil
}
