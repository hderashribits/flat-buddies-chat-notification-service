package chat

import (
	"context"
	"encoding/json"

	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/constants"
	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/model"
	"github.com/segmentio/kafka-go"
)

func SendMessage(msg model.ChatMessage) error {
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    constants.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	messageBytes, _ := json.Marshal(msg)

	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(msg.ReceiverID),
			Value: messageBytes,
		})
}
