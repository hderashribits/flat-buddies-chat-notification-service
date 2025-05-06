package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/constants"
	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/model"
	ws "github.com/hderashribits/flat-buddies-chat-notification-service/internal/websocket"
	"github.com/segmentio/kafka-go"
)

func StartChatListener() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   constants.KafkaTopic,
		GroupID: "notification-group",
	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		var notification model.Notification
		if err := json.Unmarshal(m.Value, &notification); err != nil {
			fmt.Println("Failed to parse notification:", err)
			continue
		}

		switch notification.Type {
		case "message":
			fmt.Printf("🔔 New message for %s: %s\n", notification.UserID, notification.Content)
		case "match":
			fmt.Printf("🎉 Match found for %s: %s\n", notification.UserID, notification.Content)
		default:
			fmt.Printf("ℹ️ Notification received from %s: %s\n", notification.UserID, notification.Content)
		}

		// Forward to WebSocket
		payload, _ := json.Marshal(notification)
		ws.NotificationHub.SendToUser(notification.UserID, payload)
	}
}
