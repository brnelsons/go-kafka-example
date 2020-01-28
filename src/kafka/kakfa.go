package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

// Wrapper to simplify consuming messages.
type Consumer struct {
	reader *kafka.Reader
}

// Creates a new consumer connected to kafka, ready to start consuming messages.
func NewConsumer(kafkaBrokerUrls []string, topic string, groupId string) *Consumer {
	// https://godoc.org/github.com/segmentio/kafka-go#Reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        kafkaBrokerUrls,
		Topic:          topic,
		GroupID:        groupId,
		Partition:      0,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		MaxWait:        time.Second * 10,
	})
	return &Consumer{reader: r}
}

// Consumes a single message from the kafka stream.
func (consumer *Consumer) ConsumeOne() (message kafka.Message, err error) {
	return consumer.reader.ReadMessage(context.Background())
}

// Closes the connection to kafka
func (consumer *Consumer) Close() {
	_ = consumer.reader.Close()
}
