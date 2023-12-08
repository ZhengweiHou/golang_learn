package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

const (
	kafkaBrokers    = "localhost:9092"
	kafkaTopic      = "worker_registration"
	heartbeatTopic  = "heartbeat"
	heartbeatPeriod = 15 * time.Second
)

type WorkerInfo struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func main() {
	// 从Kafka获取Worker信息
	consumer, err := sarama.NewConsumer([]string{kafkaBrokers}, nil)
	if err != nil {
		log.Fatal("Error creating consumer:", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(kafkaTopic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("Error creating partition consumer:", err)
	}
	defer partitionConsumer.Close()

	workerList := make(map[string]WorkerInfo)

	// 定期检查Worker列表和续约信息
	go func() {
		ticker := time.NewTicker(heartbeatPeriod)
		for range ticker.C {
			checkWorkers(workerList)
		}
	}()

	// 处理Worker注册和反注册信息
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var workerInfo WorkerInfo
			err := json.Unmarshal(msg.Value, &workerInfo)
			if err != nil {
				fmt.Println("Error decoding worker registration message:", err)
				continue
			}

			if msg.Topic == kafkaTopic {
				// 处理注册信息
				workerList[workerInfo.IP+":"+workerInfo.Port] = workerInfo
				fmt.Println("Worker registered:", workerInfo)
			} else if msg.Topic == heartbeatTopic {
				// 更新续约信息
				workerList[workerInfo.IP+":"+workerInfo.Port] = workerInfo
				fmt.Println("Heartbeat received from worker:", workerInfo)
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("Error receiving message:", err)
		}
	}
}

func checkWorkers(workerList map[string]WorkerInfo) {
	// TODO: 实现健康检查逻辑，处理续约信息
	// 通过负载均衡选择健康的Worker进行业务接口调用
	fmt.Println("Checking worker health and performing load balancing...")
}
