package consumer

import (
	"context"
	"log"
	"time"

	"notification-service/internal/handler"

	"github.com/segmentio/kafka-go"
)

func Start() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "chat-messages",
		GroupID:  "notification-service",
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	defer r.Close()

	log.Println("✅ Notification consumer started...")

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		msg, err := r.ReadMessage(ctx)
		cancel()

		if err != nil {
			if err == context.DeadlineExceeded {
				continue
			}

			log.Printf("❌ Kafka read error: %v", err)
			continue
		}

		// Pass message to handler
		handler.HandleMessage(msg.Value)
	}
}
