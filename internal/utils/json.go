package utils

import (
	"encoding/json"

	notification "github.com/ahyaghoubi/notification-sdk"
)

// MarshalNotification safely marshals a notification to JSON
func MarshalNotification(n *notification.Notification) ([]byte, error) {
	data, err := json.Marshal(n)
	if err != nil {
		return nil, notification.NewError(notification.Internal, "failed to marshal notification: "+err.Error())
	}
	return data, nil
}

// UnmarshalNotification safely unmarshals JSON to a notification
func UnmarshalNotification(data []byte) (*notification.Notification, error) {
	var n notification.Notification
	if err := json.Unmarshal(data, &n); err != nil {
		return nil, notification.NewError(notification.Internal, "failed to unmarshal notification: "+err.Error())
	}
	return &n, nil
}
