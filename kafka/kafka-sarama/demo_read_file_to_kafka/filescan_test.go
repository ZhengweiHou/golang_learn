package simple

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/IBM/sarama"
)

func TestFileScan(t *testing.T) {
	fileName := "data/test.txt"
	brokers := []string{"localhost:9092"}
	// topic := "test-topic"
	// topic := "test_hzw"
	topic := "hzw-transaction"

	// 打开文件
	fmt.Printf("开始测试文件 %s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("打开文件失败: %v", err)
		return
	}
	defer file.Close()

	// 计算文件的md5
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Printf("计算文件MD5失败: %v", err)
		return
	}
	fileMD5 := hex.EncodeToString(hash.Sum(nil))
	fmt.Printf("md5: %s\n", fileMD5)

	// 创建Kafka生产者
	producer, err := NewProducer(brokers)
	if err != nil {
		fmt.Printf("创建Kafka生产者失败: %v", err)
		return
	}
	defer producer.Close()

	// 重置文件指针
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Printf("重置文件指针失败: %v", err)
		return
	}

	// 扫描文件并发送到Kafka
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lineCount++

		key := fmt.Sprintf("%s-%s-%d", fileName, fileMD5, lineCount)
		headers := []sarama.RecordHeader{
			{
				Key:   []byte("file-name"),
				Value: []byte(fileName),
			},
			{
				Key:   []byte("file-md5"),
				Value: []byte(fileMD5),
			},
			{
				Key:   []byte("line-count"),
				Value: []byte(fmt.Sprintf("%d", lineCount)),
			},
		}
		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Key:       sarama.StringEncoder(key),
			Headers:   headers,
			Value:     sarama.StringEncoder(line),
			Timestamp: time.Now(),
		}

		// 发送消息到Kafka
		start := time.Now()
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Printf("发送Kafka消息失败: %v", err)
			return
		}
		fmt.Printf("消息发送成功 - 分区: %d, 偏移量: %d, 耗时: %v\n", partition, offset, time.Since(start))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("扫描文件时出错: %v", err)
	}

}
