package utils

import (
	"encoding/json"

	"github.com/getsentry/sentry-go"

	notification "github.com/ahyaghoubi/notification-sdk"
)

// MarshalNotification safely marshals a notification to JSON
func MarshalNotification(n *notification.Notification) ([]byte, error) {
	data, err := json.Marshal(n)
	if err != nil {
		errWrap := notification.NewError(notification.Internal, "failed to marshal notification: "+err.Error())
		sentry.CaptureException(errWrap)
		return nil, errWrap
	}
	return data, nil
}

// UnmarshalNotification safely unmarshals JSON to a notification
func UnmarshalNotification(data []byte) (*notification.Notification, error) {
	var n notification.Notification
	if err := json.Unmarshal(data, &n); err != nil {
		errWrap := notification.NewError(notification.Internal, "failed to unmarshal notification: "+err.Error())
		sentry.CaptureException(errWrap)
		return nil, errWrap
	}
	return &n, nil
}
