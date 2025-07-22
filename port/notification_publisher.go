package port

import "github.com/ahyaghoubi/notification-sdk/domain"

type NotificationPublisherPort interface {
	PublishNotification(clientID string, message string, notificationType domain.NotificationType, source string) error
	PublishCustomNotification(clientID string, notification *domain.Notification) error
	Close() error
}
