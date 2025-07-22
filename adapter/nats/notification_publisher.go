package natsadapter

import (
	"encoding/json"
	"time"

	"github.com/ahyaghoubi/notification-sdk/internal/domain"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type NotificationPublisher struct {
	nc            *nats.Conn
	js            nats.JetStreamContext
	subjectPrefix string
}

func NewNotificationPublisher(natsURL, subjectPrefix string) (*NotificationPublisher, error) {
	nc, err := nats.Connect(natsURL, nats.MaxReconnects(-1))
	if err != nil {
		return nil, domain.NewError(domain.Internal, "failed to connect to NATS: "+err.Error())
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, domain.NewError(domain.Internal, "failed to create JetStream context: "+err.Error())
	}

	publisher := &NotificationPublisher{
		nc:            nc,
		js:            js,
		subjectPrefix: subjectPrefix,
	}

	return publisher, nil
}

func (p *NotificationPublisher) PublishNotification(clientID string, message string, notificationType domain.NotificationType, source string) error {
	if clientID == "" {
		return domain.NewError(domain.InvalidArguments, "clientID cannot be empty")
	}
	if message == "" {
		return domain.NewError(domain.InvalidArguments, "message cannot be empty")
	}
	if source == "" {
		return domain.NewError(domain.InvalidArguments, "source cannot be empty")
	}

	notification := &domain.Notification{
		ID:        uuid.New().String(),
		ClientID:  clientID,
		Message:   message,
		Type:      notificationType,
		Read:      false,
		CreatedAt: time.Now().UTC(),
		Source:    source,
	}

	data, err := json.Marshal(notification)
	if err != nil {
		return domain.NewError(domain.Internal, "failed to marshal notification: "+err.Error())
	}

	subject := p.subjectPrefix + "." + clientID
	err = p.nc.Publish(subject, data)
	if err != nil {
		return domain.NewError(domain.Internal, "failed to publish notification: "+err.Error())
	}

	return nil
}

func (p *NotificationPublisher) PublishCustomNotification(clientID string, notification *domain.Notification) error {
	if clientID == "" {
		return domain.NewError(domain.InvalidArguments, "clientID cannot be empty")
	}
	if notification == nil {
		return domain.NewError(domain.InvalidArguments, "notification cannot be nil")
	}

	if notification.ClientID == "" {
		notification.ClientID = clientID
	}
	if notification.ID == "" {
		notification.ID = uuid.New().String()
	}
	if notification.CreatedAt.IsZero() {
		notification.CreatedAt = time.Now().UTC()
	}

	data, err := json.Marshal(notification)
	if err != nil {
		return domain.NewError(domain.Internal, "failed to marshal notification: "+err.Error())
	}

	subject := p.subjectPrefix + "." + clientID
	err = p.nc.Publish(subject, data)
	if err != nil {
		return domain.NewError(domain.Internal, "failed to publish notification: "+err.Error())
	}

	return nil
}

func (p *NotificationPublisher) Close() error {
	if p.nc != nil {
		p.nc.Close()
	}
	return nil
}
