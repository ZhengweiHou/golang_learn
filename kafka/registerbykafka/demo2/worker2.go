package main

import (
	"encoding/json"
	"fmt"
	"hzw/golang_learn/hzwutils/ip"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	regkafkaBrokers    = "localhost:9092"
	regkafkaTopic      = "_topic_hzw_dfregister"
	regheartbeatPeriod = 300 * time.Second

	// kafkaBrokers    = "localhost:9092"
	// kafkaTopic      = "worker_registration"
	// heartbeatTopic  = "heartbeat"
	// heartbeatPeriod = 5 * time.Second
)

// worker.go
var (
	workerInfoMutex sync.Mutex
	instance        *Instance
)

type Instance struct {
	AppID         string            `json:"appID"`
	Host          string            `json:"host"`
	Port          int               `json:"port"` // web 端口
	Version       string            `json:"version"`
	Metadata      map[string]string `json:"metadata"`
	Status        uint32            `json:"status"`
	UpTs          int64             `json:"upTs"`          // 启动时间
	RenewTs       int64             `json:"renewTs"`       // 续约时间
	LatestTs      int64             `json:"latestTs"`      // 节点信息更新时间
	ExpirDuration time.Duration     `json:"expirDuration"` // 过期间隔
}

func (i *Instance) Key() string {
	return fmt.Sprintf("%s_%d", i.Host, i.Port)
}

// 主程序入口
func main() {

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	port, _ := ip.GetAvailablePort("127.0.0.1", 9900)

	now := time.Now().UnixMilli()
	// 收集自身元数据信息
	instance = &Instance{
		AppID:         "hzw",
		Host:          "127.0.0.1",
		Port:          port,
		RenewTs:       now,
		ExpirDuration: time.Second * 10,
	}

	// 初始化 Kafka 生产者
	kafkaProducer := initializeKafkaProducer()
	defer kafkaProducer.Close()

	// 注册到 Kafka
	logrus.Info("注册到 Kafka")
	registerToKafka(kafkaProducer)

	// 启动定时续约任务
	go periodicRenewal(kafkaProducer)

	// 启动 REST 接口
	go startRESTServer(instance)

	logrus.Info("启动完成")
	// 等待中断信号，进行反注册
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGKILL)
	<-sigterm

	logrus.Info("worker 停止")
	// 发送反注册消息到 Kafka
	unregisterFromKafka(kafkaProducer)
}

// 初始化 Kafka 生产者
func initializeKafkaProducer() sarama.SyncProducer {
	// config := sarama.NewConfig()
	producer, err := sarama.NewSyncProducer([]string{regkafkaBrokers}, nil)
	if err != nil {
		log.Fatalln("Failed to create Kafka producer:", err)
	}
	return producer
}

// 注册到 Kafka
func registerToKafka(producer sarama.SyncProducer) {
	message, err := json.Marshal(instance)
	if err != nil {
		log.Println("Error encoding registration message:", err)
		return
	}

	// 将注册消息发送到kafka (注册消息本质上还是续约消息)
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: regkafkaTopic,
		Value: sarama.StringEncoder(message),
	})

	if err != nil {
		panic(err)
	}
}

// 启动定时续约任务
func periodicRenewal(producer sarama.SyncProducer) {
	for {
		time.Sleep(regheartbeatPeriod) // Adjust the interval as needed

		// Implement logic to renew worker registration
		workerInfoMutex.Lock()
		instance.RenewTs = time.Now().UnixMilli()
		workerInfoMutex.Unlock()

		message, err := json.Marshal(instance)
		if err != nil {
			log.Println("Error encoding renewal message:", err)
			continue
		}

		// 将续约消息发送到kafka
		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Key:   sarama.StringEncoder(instance.Key()),
			Topic: regkafkaTopic,
			Value: sarama.StringEncoder(message),
		})
		if err != nil {
			// 续约消息发送失败
			logrus.Error("续约消息发送失败")
		}
	}
}

// 启动 REST 接口
func startRESTServer(inst *Instance) {
	router := gin.Default()

	// Define REST API routes
	router.GET("/naming/check", func(c *gin.Context) {
		workerInfoMutex.Lock()
		defer workerInfoMutex.Unlock()

		c.JSON(http.StatusOK, inst)
	})

	router.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		rmsg := helloWorker(name)
		c.String(http.StatusOK, rmsg)
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name, ok := c.Params.Get("name")
		rmsg := ""
		if ok {
			rmsg = helloWorker(name)
		}
		c.String(http.StatusOK, rmsg)
	})

	// Start the server
	if err := router.Run(fmt.Sprintf(":%d", inst.Port)); err != nil {
		log.Fatalln("Failed to start REST API server:", err)
	}
}

// 发送反注册消息到 Kafka
func unregisterFromKafka(producer sarama.SyncProducer) {

	workerTemp := *instance
	workerTemp.ExpirDuration = -1 // 反注册=租期设置为-1

	message, err := json.Marshal(workerTemp)
	if err != nil {
		logrus.Error("Error encoding renewal message:", err)
	}

	// 将续约消息发送到kafka
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: regkafkaTopic,
		Value: sarama.StringEncoder(message),
	})
	if err != nil {
		// 续约消息发送失败
		logrus.Error("反注册消息发送失败")
	}
}

func helloWorker(msg string) string {
	rmsg := fmt.Sprintf("helloworker,workerport:%d, msg:%s", instance.Port, msg)
	logrus.Info(rmsg)
	return rmsg
}
