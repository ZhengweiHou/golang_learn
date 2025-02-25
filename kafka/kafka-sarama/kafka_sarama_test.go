package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/IBM/sarama"
)

const (
	// brokerList = "localhost:9092" // Kafka broker 地址
	brokerList = "localhost:9092" // Kafka broker 地址
	topic      = "hzwkfk_topic"   // Kafka 主题名称
	// groupID    = "hzwkfk-group"   // 消费者组 ID
)

func Test_produce(t *testing.T) {
	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), nil)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()
	time.Now()
	// 创建头部信息
	headers := []sarama.RecordHeader{
		{Key: []byte("auth"), Value: []byte("hzw")},
		{Key: []byte("time"), Value: []byte(time.Now().Format("2006-01-02 15:04:05"))},
	}

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := rd.Intn(10)
	// randomNumber := rand.Intn(10)
	// 要序列化的 map
	data := map[string]interface{}{
		"id":   i,
		"name": "hzw10",
		// "vlist": []string{"item1", "item2", "item3"},
	}
	// 将 map 序列化成 JSON
	jsonData, err := json.Marshal(data)

	// 生产者发送消息
	message := &sarama.ProducerMessage{
		Key: sarama.StringEncoder(fmt.Sprintf("%v", data["id"])), // 使用id作为消息的key
		// Partition: 0,                                                   // 指定消息分区
		Topic:   topic,
		Value:   sarama.StringEncoder(jsonData),
		Headers: headers,
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}

func Test_consume(t *testing.T) {
	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer(strings.Split(brokerList, ","), nil)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	// 消费者订阅主题
	// partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	// partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, 0)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	// 启动消费者协程
	go func() {
		// 处理消费者消息
		for {
			select {
			case msg := <-partitionConsumer.Messages():

				head := make(map[string]string)
				for _, v := range msg.Headers {
					head[string(v.Key)] = string(v.Value)
				}

				jsonBytes, _ := json.Marshal(head)
				log.Printf("head_json:%s\n", string(jsonBytes))

				log.Printf("Received head: %v, offset: %d, message: %s", head, msg.Offset, string(msg.Value))
			}
		}
	}()
	// go consumeMessages(partitionConsumer)

	// 等待程序退出信号
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// 等待中断信号或者消费者完成
	select {
	case <-sigchan:
		log.Println("Received an interrupt, stopping consumer.")
	case <-time.After(3 * time.Second):
		log.Println("Timeout, stopping consumer.")
	}
}

func Test_topicPartitionInfo1(t *testing.T) {
	// 设置Kafka broker地址
	brokerList := []string{"localhost:9092"}

	// 指定Kafka主题
	topic := "hzwkfk_topic"

	// 创建Kafka配置
	config := sarama.NewConfig()

	// 使用Kafka版本1.0.1
	config.Version = sarama.V1_0_0_0

	// 创建Kafka生产者
	client, err := sarama.NewClient(brokerList, config)
	if err != nil {
		log.Fatal("Error creating Kafka client: ", err)
		return
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatal("Error closing Kafka client: ", err)
		}
	}()

	// 获取指定主题的所有分区
	partitions, err := client.Partitions(topic)
	if err != nil {
		log.Fatal("Error getting partitions: ", err)
		return
	}

	// 打印每个分区的信息
	for _, partition := range partitions {
		// 获取分区的Leader
		leader, err := client.Leader(topic, partition)
		if err != nil {
			log.Fatal("Error getting leader for partition ", partition, ": ", err)
			return
		}

		// 获取分区的副本
		replicas, err := client.Replicas(topic, partition)
		if err != nil {
			log.Fatal("Error getting replicas for partition ", partition, ": ", err)
			return
		}

		// 获取分区的ISR（In-Sync Replicas）
		isr, err := client.InSyncReplicas(topic, partition)
		if err != nil {
			log.Fatal("Error getting ISR for partition ", partition, ": ", err)
			return
		}

		// 获取指定分区的当前偏移量
		offset, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatal("Error getting offset: ", err)
			return
		}

		fmt.Printf("Partition: %d, Leader: %d, Replicas: %v, ISR: %v, offset:%d\n", partition, leader.ID(), replicas, isr, offset)
	}

}

func Test_sar_consume2(t *testing.T) {
	topic := "notexisttopic"
	cs, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	ps, err := cs.Partitions(topic)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", ps)

	pc, err := cs.ConsumePartition(topic, 1, sarama.OffsetNewest)
	if err != nil {
		fmt.Println(err.Error())
	}
	select {
	case <-time.After(time.Second):
	case msg := <-pc.Messages():
		fmt.Println(msg.Value)
	}
}

func Test_produce_batch1(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = 0
	// sarama.NewClient(addrs []string, conf *sarama.Config)

	// 创建 Kafka 生产者
	// producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), nil)
	producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), config)
	if err != nil {
		log.Fatalf("Error creating producer: %s", err.Error())
	}
	defer producer.Close()

	time.Now()
	// 创建头部信息
	// headers := []sarama.RecordHeader{
	// 	{Key: []byte("auth"), Value: []byte("hzw")},
	// 	{Key: []byte("time"), Value: []byte(time.Now().Format("2006-01-02 15:04:05"))},
	// }

	// rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	msgs := make([]*sarama.ProducerMessage, 0)

	for n := 0; n < 10; n++ {
		// i := rd.Intn(10)
		i := n

		datastr := fmt.Sprintf("msg-%d", i)
		// 生产者发送消息
		msgs = append(msgs, &sarama.ProducerMessage{
			// Key:   sarama.StringEncoder(fmt.Sprintf("%d", i)), // 使用id作为消息的key
			// Key:   sarama.StringEncoder(""), // key 为空的情况
			Topic: topic,
			Value: sarama.StringEncoder(datastr),
		})
	}

	err = producer.SendMessages(msgs)

	partmsgs := make(map[string]string)

	for _, msg := range msgs {
		var k []byte
		if msg.Key != nil {
			k, _ = msg.Key.Encode()
		}
		v, _ := msg.Value.Encode()

		partmsgs[fmt.Sprintf("%s", k)] = fmt.Sprintf("%s-%s", k, v)
	}

	fmt.Printf("%v\n", partmsgs)
}
