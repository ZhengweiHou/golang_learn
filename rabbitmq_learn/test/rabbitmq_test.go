package test

import (
	"log/slog"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestRabbitmq1(t *testing.T) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ", t)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel", t)

	defer ch.Close()
	slog.Info("channel is open")

	// 获取匿名队列

	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue", t)
	slog.Info("queue is declared", "name", q.Name)

	// 写入信息
	// err = ch.Publish(
	// 	"",     // exchange
	// 	q.Name, // routing key
	// 	false,  // mandatory
	// 	false,  // immediate
	// 	amqp.Publishing{
	// 		ContentType:  "text/plain",
	// 		Body:         []byte("Hello World!"),
	// 		DeliveryMode: amqp.Persistent,
	// 	})
	failOnError(err, "Failed to publish a message", t)
	slog.Info("message is published", "queue", q.Name)

}

func failOnError(err error, msg string, t *testing.T) {
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
}
