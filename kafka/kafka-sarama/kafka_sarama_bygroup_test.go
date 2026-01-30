package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

func Test_consumeByGroup(t *testing.T) {
	brokerList := "localhost:9092" // Kafka broker 地址
	topic := "hzwkfk_topic"        // Kafka 主题名称
	groupID := "hzwkfk-group1"

	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	// sarama client
	client, err := sarama.NewClient(strings.Split(brokerList, ","), config)

	// 创建 Kafka 消费者组
	consumerGroup, err := sarama.NewConsumerGroupFromClient(groupID, client) // 通过client创建消费者组
	// consumerGroup, err := sarama.NewConsumerGroup(strings.Split(brokerList, ","), groupID, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumerGroup.Close()

	// 获取当前topic所有分区的offset
	parts, _ := client.Partitions(topic)
	topicoffsets := make(map[int32]int64)
	for _, parti := range parts {
		begainoffset, _ := client.GetOffset(topic, parti, sarama.OffsetNewest)
		topicoffsets[parti] = begainoffset
	}

	// 消费者处理函数
	handler := &messageHandler{
		partBegainOffsets: make(map[string]map[int32]int64),
	}
	handler.partBegainOffsets[topic] = topicoffsets

	// 消费者消费消息(重平衡时，Consume会停止中断)
	for {
		err := consumerGroup.Consume(context.Background(), []string{topic}, handler)
		if err != nil {
			log.Printf("Error consuming messages: %v", err)
		}
	}
}

// messageHandler 实现了 sarama.ConsumerGroupHandler 接口
type messageHandler struct {
	partBegainOffsets map[string]map[int32]int64
	// []int32
}

func (m *messageHandler) Setup(session sarama.ConsumerGroupSession) error {
	// for topic, offsets := range m.partBegainOffsets {
	// 	for part, offset := range offsets {
	// 		session.GenerationID()
	// 		logrus.Infof("%s[%d]:%d\n", topic, part, offset-1)
	// 		session.ResetOffset(topic, part, offset-1, "")
	// 	}
	// }
	fmt.Println("Setup")
	return nil
}

func (m *messageHandler) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("Cleanup")
	return nil
}

func (m *messageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Println("ConsumeClaim")
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}

			head := make(map[string]string)

			for _, v := range msg.Headers {
				head[string(v.Key)] = string(v.Value)
			}
			// 处理消费者消息
			fmt.Printf("Received message: Topic=%s, Partition=%d, Offset=%d, headers=%v, Key=%s, Value=%s\n",
				msg.Topic, msg.Partition, msg.Offset, head, string(msg.Key), string(msg.Value))

			key := string(msg.Key)

			if strings.Contains(key, "1") { // 模拟不提交偏移量
				continue
			}
			// 手动提交偏移量（消息处理完成后确认提交，）
			// 搭配 config.Consumer.Offsets.AutoCommit.Enable = false 关闭自动提交
			session.MarkMessage(msg, "") // 会最终提交偏移量，即使前面的消息处理失败，未提交偏移量
		case <-session.Context().Done():
			fmt.Println("session.Done")
			return nil
		}
	}

}
