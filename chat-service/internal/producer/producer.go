package producer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafka(brokers []string, topic string) {
	writer = &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	log.Println("✅ Kafka producer initialized")
}

func SendMessage(ctx context.Context, value []byte) error {
	err := writer.WriteMessages(ctx, kafka.Message{
		Value: value,
	})
	if err != nil {
		log.Printf("❌ Kafka send error: %v", err)
	}
	return err
}
