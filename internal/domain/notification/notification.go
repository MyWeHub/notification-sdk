package notification

import (
	"time"
)

// NotificationType defines the type of notification
type NotificationType int32

// Notification types
const (
	TypeInfo    NotificationType = 0
	TypeWarning NotificationType = 1
	TypeError   NotificationType = 2
	TypeSuccess NotificationType = 3
	TypeSystem  NotificationType = 4
)

// Notification represents a message sent to a user
type Notification struct {
	ID        string           `json:"id"`
	ClientID  string           `json:"client_id"`
	UserID    string           `json:"user_id"`
	Message   string           `json:"message"`
	Type      NotificationType `json:"type"`
	Read      bool             `json:"read"`
	CreatedAt time.Time        `json:"created_at"`
	Source    string           `json:"source"`
}

// NotificationEvent represents a notification with an event ID for SSE
type NotificationEvent struct {
	Notification *Notification
	EventID      string
}

// MarkAsRead marks the notification as read
func (n *Notification) MarkAsRead() {
	n.Read = true
}

// IsValid checks if the notification has the required fields
func (n *Notification) IsValid() bool {
	return n.UserID != "" && n.Message != "" && n.Source != ""
}
