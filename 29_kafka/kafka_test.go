package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
)

func Test_kafka(t *testing.T) {
	topic := "hzwkfk"
	brokers := make([]string, 0)
	brokers = append(brokers, "localhost:9092")
	go func() {
		tick := time.NewTicker(time.Second * 5)
		for {
			<-tick.C
			produceMessage(topic, brokers, fmt.Sprintf("time:%v,msg:%v", time.Now(), "hello"))
		}
	}()

	consumeMessages(topic, brokers, "")
}

func produceMessage(topic string, brokers []string, message string) {
	// 创建一个生产者
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	// 创建消息
	msg := kafka.Message{
		Key:   []byte(nil),
		Value: []byte(message),
	}

	// 发送消息
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Fatal("无法发送消息到Kafka:", err)
	}

	// 关闭生产者连接
	err = writer.Close()
	if err != nil {
		log.Fatal("无法关闭Kafka生产者:", err)
	}
}

func consumeMessages(topic string, brokers []string, groupID string) {
	// 创建一个消费者
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		// GroupID: groupID,
	})

	// 循环接收消息
	for {
		// 读取消息
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("无法从Kafka读取消息:", err)
		}

		// 处理消息
		fmt.Printf("接收到消息: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
	}

	// 关闭消费者连接
	err := reader.Close()
	if err != nil {
		log.Fatal("无法关闭Kafka消费者:", err)
	}
}
