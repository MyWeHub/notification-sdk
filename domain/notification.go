package domain

import (
	"time"
)

type NotificationType int32

const (
	TypeInfo    NotificationType = 0
	TypeWarning NotificationType = 1
	TypeError   NotificationType = 2
	TypeSuccess NotificationType = 3
	TypeSystem  NotificationType = 4
)

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

type NotificationEvent struct {
	Notification *Notification
	EventID      string
}

func (n *Notification) MarkAsRead() {
	n.Read = true
}

func (n *Notification) IsValid() bool {
	return n.UserID != "" && n.Message != "" && n.Source != ""
}
