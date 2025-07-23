package notification

// PublisherPort defines the interface for publishing notifications
type PublisherPort interface {
	PublishNotification(clientID string, message string, notificationType NotificationType, source string) error
	PublishCustomNotification(clientID string, notification *Notification) error
	Close() error
}
