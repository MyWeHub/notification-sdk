package nats

import (
	"github.com/ahyaghoubi/notification-sdk/internal/natsutil"
	"github.com/ahyaghoubi/notification-sdk/internal/utils"
	"github.com/ahyaghoubi/notification-sdk/internal/validation"

	"github.com/getsentry/sentry-go"

	notification "github.com/ahyaghoubi/notification-sdk"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type Publisher struct {
	nc            *nats.Conn
	js            nats.JetStreamContext
	subjectPrefix string
}

// NewPublisher creates a new NATS notification publisher with default options
func NewPublisher(natsURL, subjectPrefix string) (*Publisher, error) {
	nc, err := natsutil.ConnectWithRetry(natsURL, 3)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	js, err := natsutil.CreateJetStreamContext(nc)
	if err != nil {
		nc.Close() // Clean up connection on error
		sentry.CaptureException(err)
		return nil, err
	}

	return &Publisher{
		nc:            nc,
		js:            js,
		subjectPrefix: subjectPrefix,
	}, nil
}

// NewPublisherWithOptions creates a new NATS notification publisher with custom options
func NewPublisherWithOptions(natsURL, subjectPrefix string, opts ...nats.Option) (*Publisher, error) {
	nc, err := natsutil.ConnectWithCustomOptions(natsURL, opts...)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	js, err := natsutil.CreateJetStreamContext(nc)
	if err != nil {
		nc.Close()
		sentry.CaptureException(err)
		return nil, err
	}

	return &Publisher{
		nc:            nc,
		js:            js,
		subjectPrefix: subjectPrefix,
	}, nil
}

func (p *Publisher) PublishNotification(clientID string, message string, notificationType notification.NotificationType, source string) error {
	// Use internal validation
	if err := validation.ValidateClientID(clientID); err != nil {
		return err
	}
	if err := validation.ValidateMessage(message); err != nil {
		return err
	}
	if err := validation.ValidateSource(source); err != nil {
		return err
	}

	notif := &notification.Notification{
		ID:        uuid.New().String(),
		ClientID:  clientID,
		Message:   message,
		Type:      notificationType,
		Read:      false,
		CreatedAt: utils.UTCNow(),
		Source:    source,
	}

	return p.publishNotification(notif)
}

func (p *Publisher) PublishCustomNotification(clientID string, notif *notification.Notification) error {
	if err := validation.ValidateClientID(clientID); err != nil {
		return err
	}
	if err := validation.ValidateNotification(notif); err != nil {
		return err
	}

	// Auto-fill missing fields
	if notif.ClientID == "" {
		notif.ClientID = clientID
	}
	if notif.ID == "" {
		notif.ID = uuid.New().String()
	}
	if utils.IsZeroTime(notif.CreatedAt) {
		notif.CreatedAt = utils.UTCNow()
	}

	return p.publishNotification(notif)
}

// publishNotification is a private helper method that handles the actual publishing
func (p *Publisher) publishNotification(notif *notification.Notification) error {
	// Use internal JSON utility
	data, err := utils.MarshalNotification(notif)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Use internal subject builder
	subject := natsutil.BuildSubject(p.subjectPrefix, notif.ClientID)

	// Validate subject before publishing
	if !natsutil.ValidateSubject(subject) {
		err := notification.NewError(notification.InvalidArguments, "invalid subject: "+subject)
		sentry.CaptureException(err)
		return err
	}

	err = p.nc.Publish(subject, data)
	if err != nil {
		err := notification.NewError(notification.Internal, "failed to publish notification: "+err.Error())
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (p *Publisher) Close() error {
	if p.nc != nil {
		p.nc.Close()
	}
	return nil
}

// IsConnected returns true if the NATS connection is active
func (p *Publisher) IsConnected() bool {
	return p.nc != nil && p.nc.IsConnected()
}

// GetConnectionStatus returns the current connection status
func (p *Publisher) GetConnectionStatus() nats.Status {
	if p.nc == nil {
		return nats.CLOSED
	}
	return p.nc.Status()
}
