package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"chat-service/internal/models"
	"chat-service/internal/producer"
)

func HandleMatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var match struct {
		User1ID string `json:"user1_id"`
		User2ID string `json:"user2_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	notificationMessage := models.ChatMessage{
		SenderID:   match.User1ID,
		ReceiverID: match.User2ID,
		Content:    "You are matched with " + match.User1ID,
		Timestamp:  time.Now().Unix(),
	}

	payload, err := json.Marshal(notificationMessage)
	if err != nil {
		http.Error(w, "Failed to encode notification", http.StatusInternalServerError)
		return
	}

	err = producer.SendMessage(r.Context(), payload)
	if err != nil {
		http.Error(w, "Failed to send notification to Kafka", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Match notification sent to Kafka",
		"user1":   match.User1ID,
		"user2":   match.User2ID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(response)
}
