package handler

import (
	"encoding/json"
	"log"

	"notification-service/internal/models"
)

func HandleMessage(value []byte) {
	// Check if it's valid JSON
	if !json.Valid(value) {
		log.Printf("❌ Invalid JSON: %s", string(value))
		return
	}

	var msg models.NotificationMessage
	err := json.Unmarshal(value, &msg)
	if err != nil {
		log.Printf("❌ Failed to decode message: %v", err)
		return
	}

	// Sanity check: log if sender/receiver is empty
	if msg.SenderID == "" || msg.ReceiverID == "" {
		log.Printf("⚠️ Possibly non-chat message: %s", string(value))
		return
	}

	log.Printf("📬 New chat from %s to %s: %s", msg.SenderID, msg.ReceiverID, msg.Content)
}
