package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
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

	// 消费者消费消息
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
	for topic, offsets := range m.partBegainOffsets {
		for part, offset := range offsets {
			session.GenerationID()
			logrus.Infof("%s[%d]:%d\n", topic, part, offset-1)
			session.ResetOffset(topic, part, offset-1, "")
		}
	}
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
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}

			head := make(map[string]string)

			for _, v := range message.Headers {
				head[string(v.Key)] = string(v.Value)
			}
			// 处理消费者消息
			fmt.Printf("Received message: Topic=%s, Partition=%d, Offset=%d, headers=%v, Key=%s, Value=%s\n",
				message.Topic, message.Partition, message.Offset, head, string(message.Key), string(message.Value))
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			fmt.Println("session.Done")
			return nil
		}
	}

}
