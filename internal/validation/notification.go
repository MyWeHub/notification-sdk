package validation

import (
	"strings"

	notification "github.com/ahyaghoubi/notification-sdk"
)

// ValidateClientID checks if a client ID is valid
func ValidateClientID(clientID string) error {
	if clientID == "" {
		err := notification.NewError(notification.InvalidArguments, "clientID cannot be empty")
		return err
	}

	if len(clientID) > 255 {
		err := notification.NewError(notification.InvalidArguments, "clientID cannot exceed 255 characters")
		return err
	}

	if strings.Contains(clientID, " ") {
		err := notification.NewError(notification.InvalidArguments, "clientID cannot contain spaces")
		return err
	}

	return nil
}

// ValidateMessage checks if a message is valid
func ValidateMessage(message string) error {
	if message == "" {
		err := notification.NewError(notification.InvalidArguments, "message cannot be empty")
		return err
	}

	if len(message) > 10000 {
		err := notification.NewError(notification.InvalidArguments, "message cannot exceed 10000 characters")
		return err
	}

	return nil
}

// ValidateSource checks if a source is valid
func ValidateSource(source string) error {
	if source == "" {
		err := notification.NewError(notification.InvalidArguments, "source cannot be empty")
		return err
	}

	if len(source) > 100 {
		err := notification.NewError(notification.InvalidArguments, "source cannot exceed 100 characters")
		return err
	}

	return nil
}

// ValidateNotification performs comprehensive validation on a notification
func ValidateNotification(n *notification.Notification) error {
	if n == nil {
		err := notification.NewError(notification.InvalidArguments, "notification cannot be nil")
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
