package main

import (
	"encoding/json"
	"fmt"
	"hzw/golang_learn/hzwutils/ip"
	"log"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

const (
	kafkaBrokers = "localhost:9092"
	kafkaTopic   = "worker_registration"
)

var workerInfo *WorkerInfo

// WorkerInfo worker节点信息
type WorkerInfo struct {
	IP   string
	Port int
}

// RegistrationMessage worker的注册信息
type RegistrationMessage struct {
	WorkerInfo       *WorkerInfo   `json:"worker_info"`
	RegistrationTime time.Time     `json:"registration_time"`
	ExpirationPeriod time.Duration `json:"expiration_period"`
}

func main() {

	workerInfo = &WorkerInfo{
		Port: 9000,
	}

	ip.InitIPRanges("192.168.105")
	hostip := ip.GetLocalIP()
	port, err := ip.GetAvailablePort(hostip, workerInfo.Port)
	if err != nil {
		log.Panic(err)
	}

	workerInfo.IP = hostip
	workerInfo.Port = port

	// TODO  TEST
	wijson, _ := json.Marshal(workerInfo)
	fmt.Printf("workerInfo:%s", wijson)

	// 启动Gin服务
	router := gin.Default()

	// 创建第一个RouterGroup
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/hello/:name", helloHandler)
		apiV1.GET("/hello2/:name", helloHandler)
	}

	// 创建第二个RouterGroup
	// apiV2 := router.Group("/api/v2")
	// {
	// 	apiV2.GET("/greet/:name", greetHandler)
	// }

	go func() {
		// 定时发送续约信息
		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {
			sendHeartbeat()
		}
	}()

	// 注册到Kafka
	registerWorker()

	router.Run(fmt.Sprintf("%s:%d", workerInfo.IP, workerInfo.Port))
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func registerWorker() {
	producer, err := sarama.NewSyncProducer([]string{kafkaBrokers}, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	registrationMessage := RegistrationMessage{
		WorkerInfo:       workerInfo,
		RegistrationTime: time.Now(),
		ExpirationPeriod: 15 * time.Second, // Set expiration period as needed
	}

	message, err := json.Marshal(registrationMessage)
	if err != nil {
		panic(err)
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: kafkaTopic,
		Value: sarama.StringEncoder(message),
	})

	if err != nil {
		panic(err)
	}
}

func sendHeartbeat() {
	producer, err := sarama.NewSyncProducer([]string{kafkaBrokers}, nil)
	if err != nil {
		fmt.Println("Error creating producer:", err)
		return
	}
	defer producer.Close()

	registrationMessage := RegistrationMessage{
		WorkerInfo:       workerInfo,
		RegistrationTime: time.Now(),
		ExpirationPeriod: 15 * time.Second, // Set expiration period as needed
	}

	message, err := json.Marshal(registrationMessage)
	if err != nil {
		fmt.Println("Error encoding heartbeat message:", err)
		return
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: kafkaTopic,
		Value: sarama.StringEncoder(message),
	})

	if err != nil {
		fmt.Println("Error sending heartbeat:", err)
		return
	}

	fmt.Println("Heartbeat sent.")
}

func helloHandler(c *gin.Context) {
	// 获取URL参数
	name := c.Param("name")

	// 返回Hello + 参数
	c.String(http.StatusOK, "Hello %s", name)
}
