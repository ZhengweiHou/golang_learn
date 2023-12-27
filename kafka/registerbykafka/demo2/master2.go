package main

import (
	"encoding/json"
	"fmt"
	"hzw/golang_learn/hzwutils/ip"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const ()

type WorkerInfo struct {
	IP               string        `json:"ip"`
	Port             int           `json:"port"`
	LastRenewal      time.Time     `json:"lastRenewal"`
	ExpirationPeriod time.Duration `json:"expirationPeriod"`
}

var (
	workerList map[string]*WorkerInfo
	// workerKeys      []string
	workerIndex     int64 = 0
	workerListMutex sync.Mutex
	masterPort      int
)

// 主程序入口
func main() {
	kafkaBrokers := "localhost:9092"
	kafkaTopic := "worker_registration"
	heartbeatTopic := "heartbeat"
	heartbeatPeriod := 5 * time.Second

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logrus.Infoln("master start")
	masterPort, _ = ip.GetAvailablePort("127.0.0.1", 9901)

	// Initialize Kafka consumer
	kafkaConsumer := initializeKafkaConsumer()

	// Initialize local worker list
	workerList = make(map[string]*WorkerInfo)

	// Start periodic check task
	go periodicCheck()

	// Start REST API server
	go startMasterServer()

	// Handle Kafka messages
	go func() {
		// cp, err := kafkaConsumer.ConsumePartition(kafkaTopic, 0, sarama.OffsetOldest)
		cp, err := kafkaConsumer.ConsumePartition(kafkaTopic, 0, sarama.OffsetNewest)
		if err != nil {
			log.Fatalln("Failed to create Kafka register consumer:", err)
		}
		for {
			select {
			case msg := <-cp.Messages():
				handleKafkaMessage(msg)
			case err := <-cp.Errors():
				fmt.Println("Error receiving message:", err)
				time.Sleep(5 * time.Second)
			}
		}
	}()

	// Wait for interrupt signal
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGKILL)
	<-sigterm

	// Cleanup resources
	if err := kafkaConsumer.Close(); err != nil {
		log.Fatalln("Failed to close Kafka consumer:", err)
	}
}

// Initialize Kafka consumer
func initializeKafkaConsumer() sarama.Consumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to create Kafka consumer:", err)
	}

	return consumer
}

// Periodic check of worker list
func periodicCheck() {
	for {
		time.Sleep(5 * time.Second) // Adjust the interval as needed

		workerListMutex.Lock()
		needCheckWorkers := make([]*WorkerInfo, 0)
		for _, worker := range workerList {
			// 租约到期的worker，手动检查
			if time.Now().Sub(worker.LastRenewal) > worker.ExpirationPeriod {
				needCheckWorkers = append(needCheckWorkers, worker)
			}
		}
		workerListMutex.Unlock()
		logrus.Infof("needCheckWorker size:%d", len(needCheckWorkers))
		for _, worker := range needCheckWorkers {
			// 调用 worker的接口主动查询状态
			checkWorkerStatus(worker)
		}
	}
}

// Handle Kafka message
func handleKafkaMessage(msg *sarama.ConsumerMessage) {
	var workerInfo WorkerInfo
	err := json.Unmarshal(msg.Value, &workerInfo)
	if err != nil {
		logrus.Errorln("Error decoding Kafka message:", err)
		return
	}
	logrus.Debugf("handler register message:%v", workerInfo)

	handlerWorkInfo(&workerInfo)

}

func handlerWorkInfo(workerInfo *WorkerInfo) {
	workerListMutex.Lock()
	defer workerListMutex.Unlock()

	if time.Now().Sub(workerInfo.LastRenewal) > workerInfo.ExpirationPeriod || workerInfo.ExpirationPeriod < 0 {
		// 续约超时 或 注销(续约期限为-1) 移除worker节点
		_, ok := workerList[workerKey(workerInfo)]
		if ok {
			logrus.Infof("Delete worker: %v\n", workerInfo)
			delete(workerList, workerKey(workerInfo))
		}
	} else {
		// 节点在有效期内，续约
		workerList[workerKey(workerInfo)] = workerInfo
		logrus.Infof("Updated worker: %v\n", workerInfo)
	}
}

// Generate a unique key for a worker based on IP and Port
func workerKey(workerInfo *WorkerInfo) string {
	return fmt.Sprintf("%s:%d", workerInfo.IP, workerInfo.Port)
}

// Start REST API server
func startMasterServer() {
	router := gin.Default()

	// Define REST API routes
	router.GET("/workers", func(c *gin.Context) {
		workerListMutex.Lock()
		defer workerListMutex.Unlock()

		c.JSON(http.StatusOK, workerList)
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name, ok := c.Params.Get("name")
		rmsg := ""
		if ok {
			rmsg = helloMaster(name)

		}

		c.String(http.StatusOK, rmsg)
	})

	router.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		rmsg := helloMaster(name)
		c.String(http.StatusOK, rmsg)
	})

	// Start the server
	// if err := router.Run(":9001"); err != nil {
	if err := router.Run(fmt.Sprintf(":%d", masterPort)); err != nil {
		log.Fatalln("Failed to start REST API server:", err)
	}
}

func checkWorkerStatus(worker *WorkerInfo) {
	// 构建 Worker 的 /info 接口地址
	workerInfoURL := fmt.Sprintf("http://%s:%d/info", worker.IP, worker.Port)

	var workerInfo WorkerInfo
	// 发送 HTTP GET 请求
	response, err := http.Get(workerInfoURL)
	if err != nil {
		log.Println("Error checking worker status:", err)
		// 调用异常，作为失败处理
		workerInfo = *worker
		workerInfo.ExpirationPeriod = -1
		handlerWorkInfo(&workerInfo)
		return
	}
	defer response.Body.Close()

	// 检查响应状态码
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&workerInfo)
		if err != nil {
			logrus.Errorf("Worker status check, Error parsing workerInfo result, err:%s", err.Error())
			workerInfo = *worker
			workerInfo.ExpirationPeriod = -1
		}
	} else {
		// 其他返回码，均作为状态异常，需将节点剔除
		log.Printf("Unexpected status code from worker: %d\n", response.StatusCode)
		workerInfo = *worker
		workerInfo.ExpirationPeriod = -1
	}

	handlerWorkInfo(&workerInfo)
}

func helloMaster(msg string) string {
	logrus.Infof("hello:%s", msg)
	worker := selectWorker()
	// TODO worker 请求失败，需要检查worker的可用状态，同时系统应该能处理故障转移

	workerHelloURL := fmt.Sprintf("http://%s:%d/hello/%s", worker.IP, worker.Port, msg)
	response, err := http.Get(workerHelloURL)
	if err != nil {
		logrus.Error("Error checking worker status:", err)
		return ""
	}
	defer response.Body.Close()

	repBody, _ := io.ReadAll(response.Body)
	return fmt.Sprintf("%s", repBody)

}

func selectWorker() *WorkerInfo {
	workerListMutex.Lock()
	defer workerListMutex.Unlock()
	keys := make([]string, 0, len(workerList))

	for key := range workerList {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	defer atomic.AddInt64(&workerIndex, 1)
	index := workerIndex % int64(len(keys))
	return workerList[keys[index]]
}
