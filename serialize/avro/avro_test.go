package avro

import (
	"fmt"
	"testing"

	"github.com/linkedin/goavro"
	log "github.com/sirupsen/logrus"
)

func Test_avro1(t *testing.T) {
	data, err := io.ReadFile("/home/houzw/temp/output.dat")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	log.Printf("data: %s", string(data))

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

	// 反序列化Avro数据
	native, _, err := avroCodec.NativeFromBinary(data)
	if err != nil {
		log.Println("Avro deserialization failed:", err)
	}
	record, err := avroCodec.TextualFromNative(nil, native)
	if err != nil {
		log.Println("Avro deserialization failed:", err)
	}

	fmt.Printf("Read data from file: %v\n", record)
}

func Test_avro2(t *testing.T) {

}
