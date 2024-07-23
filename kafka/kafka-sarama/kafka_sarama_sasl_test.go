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
	"github.com/xdg-go/scram"

	"github.com/IBM/sarama"
)

func Test_produceSASL(t *testing.T) {
	conf := sarama.NewConfig()

	conf.Producer.Return.Successes = true
	conf.Net.SASL.Enable = true
	// conf.Net.SASL.AuthIdentity = "SASL"
	// conf.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
	conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
		fmt.Println("new SCRAMClient")
		return &scramClient{}
		// return &XDGSCRAMClient{
		// 	HashGeneratorFcn: sha256.New,
		// }
	}

	conf.Net.SASL.Version = sarama.SASLHandshakeV1
	// conf.ApiVersionsRequest = true
	// conf.Net.SASL.Version = sarama.SASLHandshakeV0

	// conf.Net.SASL.User = "admin"
	// conf.Net.SASL.Password = "admin-secret"
	conf.Net.SASL.User = "hzw"
	conf.Net.SASL.Password = "hzw"
	// conf.Net.SASL.Password = "17aadef8b1390bb41879cc76f9ec11a1ca3b73ac84c23e377b34aa950751cf04"

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), conf)
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

func Test_consumeSASL(t *testing.T) {
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

type XDGSCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

func (x *XDGSCRAMClient) Begin(userName, password, authzID string) (err error) {
	x.Client, err = x.HashGeneratorFcn.NewClient(userName, password, authzID)
	if err != nil {
		return err
	}
	x.ClientConversation = x.Client.NewConversation()
	return nil
}

func (x *XDGSCRAMClient) Step(challenge string) (response string, err error) {
	response, err = x.ClientConversation.Step(challenge)
	return
}

func (x *XDGSCRAMClient) Done() bool {
	return x.ClientConversation.Done()
}

type scramClient struct {
	*scram.Client
	*scram.ClientConversation
}

func (x *scramClient) Begin(userName, password, authzID string) (err error) {
	x.Client, err = scram.SHA256.NewClient(userName, password, authzID)
	if err != nil {
		return err
	}
	x.ClientConversation = x.Client.NewConversation()
	return nil
}

func (x *scramClient) Step(challenge string) (string, error) {
	return x.ClientConversation.Step(challenge)
}

func (x *scramClient) Done() bool {
	return x.ClientConversation.Done()
}
