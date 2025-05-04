package notification

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/constants"
	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/model"
	"github.com/segmentio/kafka-go"
)

func SendNotification(n model.Notification) error {
	// Set the timestamp here
	n.Timestamp = time.Now().Unix()

	// Marshal the struct into JSON bytes
	value, err := json.Marshal(n)
	if err != nil {
		return err
	}

	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    constants.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(n.UserID),
			Value: value,
		})
}
