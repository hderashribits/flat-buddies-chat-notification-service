package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/chatandnotification"
	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/constants"
	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/model"
	"github.com/hderashribits/flat-buddies-chat-notification-service/internal/notification"
	ws "github.com/hderashribits/flat-buddies-chat-notification-service/internal/websocket"
)

func main() {

	// Chat
	go notification.StartChatListener()

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		var msg model.ChatMessage
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		msg.Timestamp = time.Now().Unix()

		notif := model.Notification{
			UserID:  msg.ReceiverID,
			Type:    "message",
			Content: fmt.Sprintf("Message from %s: %s", msg.SenderID, msg.Content),
		}

		if err := chatandnotification.SendMessage(notif); err != nil {
			http.Error(w, "failed to send message chat", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Message sent and chat triggered!")
	})

	// Flatmates Match
	http.HandleFunc("/match", func(w http.ResponseWriter, r *http.Request) {
		var notif model.Notification
		if err := json.NewDecoder(r.Body).Decode(&notif); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		notif.Type = "match"
		if err := chatandnotification.SendMessage(notif); err != nil {
			http.Error(w, "failed to send match notification", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Match notification sent!")
	})

	// WebSocket
	var upgrader = websocket.Upgrader{}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "user_id required", http.StatusBadRequest)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "websocket upgrade failed", http.StatusInternalServerError)
			return
		}

		ws.NotificationHub.Register(userID, conn)
		fmt.Println("🔌 WebSocket connected:", userID)
	})

	fmt.Printf("🚀 Server running on :%s (Chat API + Notification Listener)", constants.APIPort)
	if err := http.ListenAndServe(":"+constants.APIPort, nil); err != nil {
		panic(err)
	}
}
