package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/IBM/sarama"
)

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

	for n := 0; n < 20; n++ {
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
