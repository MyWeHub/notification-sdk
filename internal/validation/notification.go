package validation

import (
	"strings"

	notification "github.com/ahyaghoubi/notification-sdk"
)

// ValidateClientID checks if a client ID is valid
func ValidateClientID(clientID string) error {
	if clientID == "" {
		return notification.NewError(notification.InvalidArguments, "clientID cannot be empty")
	}

	if len(clientID) > 255 {
		return notification.NewError(notification.InvalidArguments, "clientID cannot exceed 255 characters")
	}

	if strings.Contains(clientID, " ") {
		return notification.NewError(notification.InvalidArguments, "clientID cannot contain spaces")
	}

	return nil
}

// ValidateMessage checks if a message is valid
func ValidateMessage(message string) error {
	if message == "" {
		return notification.NewError(notification.InvalidArguments, "message cannot be empty")
	}

	if len(message) > 10000 {
		return notification.NewError(notification.InvalidArguments, "message cannot exceed 10000 characters")
	}

	return nil
}

// ValidateSource checks if a source is valid
func ValidateSource(source string) error {
	if source == "" {
		return notification.NewError(notification.InvalidArguments, "source cannot be empty")
	}

	if len(source) > 100 {
		return notification.NewError(notification.InvalidArguments, "source cannot exceed 100 characters")
	}

	return nil
}

// ValidateNotification performs comprehensive validation on a notification
func ValidateNotification(n *notification.Notification) error {
	if n == nil {
		return notification.NewError(notification.InvalidArguments, "notification cannot be nil")
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
