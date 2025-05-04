package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[string]*websocket.Conn
	mu      sync.Mutex
}

var NotificationHub = Hub{
	clients: make(map[string]*websocket.Conn),
}

func (h *Hub) Register(userID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[userID] = conn
}

func (h *Hub) SendToUser(userID string, message []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if conn, ok := h.clients[userID]; ok {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
