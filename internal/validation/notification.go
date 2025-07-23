package validation

import (
	"strings"

	"github.com/getsentry/sentry-go"

	notification "github.com/ahyaghoubi/notification-sdk"
)

// ValidateClientID checks if a client ID is valid
func ValidateClientID(clientID string) error {
	if clientID == "" {
		err := notification.NewError(notification.InvalidArguments, "clientID cannot be empty")
		sentry.CaptureException(err)
		return err
	}

	if len(clientID) > 255 {
		err := notification.NewError(notification.InvalidArguments, "clientID cannot exceed 255 characters")
		sentry.CaptureException(err)
		return err
	}

	if strings.Contains(clientID, " ") {
		err := notification.NewError(notification.InvalidArguments, "clientID cannot contain spaces")
		sentry.CaptureException(err)
		return err
	}

	return nil
}

// ValidateMessage checks if a message is valid
func ValidateMessage(message string) error {
	if message == "" {
		err := notification.NewError(notification.InvalidArguments, "message cannot be empty")
		sentry.CaptureException(err)
		return err
	}

	if len(message) > 10000 {
		err := notification.NewError(notification.InvalidArguments, "message cannot exceed 10000 characters")
		sentry.CaptureException(err)
		return err
	}

	return nil
}

// ValidateSource checks if a source is valid
func ValidateSource(source string) error {
	if source == "" {
		err := notification.NewError(notification.InvalidArguments, "source cannot be empty")
		sentry.CaptureException(err)
		return err
	}

	if len(source) > 100 {
		err := notification.NewError(notification.InvalidArguments, "source cannot exceed 100 characters")
		sentry.CaptureException(err)
		return err
	}

	return nil
}

// ValidateNotification performs comprehensive validation on a notification
func ValidateNotification(n *notification.Notification) error {
	if n == nil {
		err := notification.NewError(notification.InvalidArguments, "notification cannot be nil")
		sentry.CaptureException(err)
		return err
	}

	if err := ValidateClientID(n.ClientID); err != nil {
		return err
	}

	if err := ValidateMessage(n.Message); err != nil {
		return err
	}

	if err := ValidateSource(n.Source); err != nil {
		return err
	}

	return nil
}
