package model

type ChatMessage struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
	Timestamp  int64  `json:"timestamp"`
}

type Notification struct {
	UserID    string `json:"user_id"`
	Type      string `json:"type"`    // "message" or "match"
	Content   string `json:"content"` // main message
	Timestamp int64  `json:"timestamp"`
}
