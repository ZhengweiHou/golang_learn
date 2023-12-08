package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

const (
	// brokerList = "localhost:9092" // Kafka broker 地址
	// topic      = "hzwkfk_sarama"  // Kafka 主题名称
	groupID = "hzwkfk-group2" // 消费者组 ID
)

func Test_consumeByGroup(t *testing.T) {

	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	// config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	// config.Consumer.Offsets.Initial = -1

	// 创建 Kafka 消费者组
	consumerGroup, err := sarama.NewConsumerGroup(strings.Split(brokerList, ","), groupID, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumerGroup.Close()

	// 消费者处理函数
	handler := &messageHandler{}

	// 消费者消费消息
	for {
		err := consumerGroup.Consume(context.Background(), []string{topic}, handler)
		if err != nil {
			log.Printf("Error consuming messages: %v", err)
		}
	}
}

// messageHandler 实现了 sarama.ConsumerGroupHandler 接口
type messageHandler struct{}

func (m *messageHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (m *messageHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (m *messageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {

		head := make(map[string]string)

		for _, v := range message.Headers {
			head[string(v.Key)] = string(v.Value)
		}

		// 处理消费者消息
		fmt.Printf("Received message: Topic=%s, Partition=%d, Offset=%d, headers=%v, Key=%s, Value=%s\n",
			message.Topic, message.Partition, message.Offset, head, string(message.Key), string(message.Value))

		session.MarkMessage(message, "")
	}

	return nil
}
