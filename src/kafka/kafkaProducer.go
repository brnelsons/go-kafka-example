package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

// Wrapper to simplify producing messages.
type Producer struct {
	writer *kafka.Writer
}

// Creates a new producer connected to kafka, ready to start producing messages.
func NewProducer(kafkaBrokerUrls []string, topic string) *Producer {
	// https://godoc.org/github.com/segmentio/kafka-go#Reader
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      kafkaBrokerUrls,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		MaxAttempts:  3,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	})
	return &Producer{writer: writer}
}

// Consumes a single message from the kafka stream.
func (producer *Producer) Produce(value []byte) error {
	return producer.writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: value,
		},
	)
}
