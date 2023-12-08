package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/IBM/sarama"
	// "github.com/linkedin/goavro/v2"
	"github.com/linkedin/goavro"
)

const (
	// brokerList = "localhost:9092" // Kafka broker 地址
	topicAvro = "hzw-avro2" // Kafka 主题名称
	// groupID    = "hzwkfk-group"   // 消费者组 ID
)

func Test_produce_avro(t *testing.T) {

	// Avro 模式定义，这里使用字符串表示 Avro 模式
	avroSchema := `
	{
		"namespace": "com.hzw.learn.avro.HzwAcroBean",
		"type": "record",
		"name": "HzwAvroBean",
		"fields": [
			{"name": "id", "type": "string"},
			{"name": "name", "type": "string"}
		]
	} 
	`
	// avroSchema := `
	// 	{"type":"record",
	// 	"name":"KsqlDataSourceSchema",
	// 	"namespace":"io.confluent.ksql.avro_schemas",
	// 	"fields":[{"name":"id","type":["null","string"],"default":null},{"name":"name","type":["null","string"],"default":null}],
	// 	"connect.name":"io.confluent.ksql.avro_schemas.KsqlDataSourceSchema"}
	// `

	// 创建 Avro 编码器
	avroCodec, err := goavro.NewCodec(avroSchema)
	if err != nil {
		fmt.Println("Error creating Avro codec: ", err)
		return
	}

	// 配置 Kafka 生产者
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	// 创建头部信息
	// headers := []sarama.RecordHeader{
	// 	{Key: []byte("auth"), Value: []byte("hzw")},
	// 	{Key: []byte("time"), Value: []byte(time.Now().Format("2006-01-02 15:04:05"))},
	// }

	// 要发送的消息
	message := map[string]interface{}{
		"id":   "id1",
		"name": "hzw1",
		// "vlist": []string{"item1", "item2", "item3"},
	}

	// 将消息编码为 Avro 格式
	avroMessage, err := avroCodec.BinaryFromNative(nil, message)
	log.Printf("go avro date len:%d,message: %s", len(avroMessage), string(avroMessage))
	if err != nil {
		fmt.Println("Error encoding message to Avro: ", err)
		return
	}

	// // 创建 Kafka 消息
	// kafkaMessage := &sarama.ProducerMessage{
	// 	Topic:   topicAvro,                                              // 替换为你的 Kafka 主题
	// 	Key:     sarama.StringEncoder(fmt.Sprintf("%v", message["id"])), // 使用id作为消息的key
	// 	Value:   sarama.ByteEncoder(avroMessage),
	// 	Headers: headers,
	// }

	// partition, offset, err := producer.SendMessage(kafkaMessage)
	// if err != nil {
	// 	log.Fatalf("Failed to send message: %v", err)
	// }

	// fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}

func Test_consume_avro(t *testing.T) {

	// Avro 模式定义，这里使用字符串表示 Avro 模式
	avroSchema := `
	{
		"namespace": "com.hzw.learn.avro.HzwAcroBean",
		"type": "record",
		"name": "HzwAvroBean",
		"fields": [
			{"name": "id", "type": "string"},
			{"name": "name", "type": "string"}
		]
	} 
	`
	// 创建 Avro 编码器
	avroCodec, err := goavro.NewCodec(avroSchema)
	if err != nil {
		fmt.Println("Error creating Avro codec: ", err)
		return
	}

	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer(strings.Split(brokerList, ","), nil)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	// 消费者订阅主题
	// partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	// partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	partitionConsumer, err := consumer.ConsumePartition(topicAvro, 0, 0)
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
				valueData := msg.Value
				log.Printf("kfk:len:%d,message: %s", len(valueData), string(valueData))
				valueData = valueData[5:]
				log.Printf("sub:len:%d,message: %s", len(valueData), string(valueData))
				// 反序列化Avro数据
				native, _, err := avroCodec.NativeFromBinary(valueData)
				if err != nil {
					log.Println("Avro deserialization failed:", err)
					continue
				}
				record, err := avroCodec.TextualFromNative(nil, native)
				if err != nil {
					log.Println("Avro deserialization failed:", err)
					continue
				}

				log.Printf("Received head: %v, offset: %d, message: %s", head, msg.Offset, string(record))
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
