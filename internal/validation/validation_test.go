package validation

import (
	"strings"
	"testing"

	notification "github.com/MyWeHub/notification-sdk"
)

func TestValidateClientID(t *testing.T) {
	tests := []struct {
		name     string
		clientID string
		wantErr  bool
	}{
		{"valid client ID", "client-123", false},
		{"empty client ID", "", true},
		{"client ID with spaces", "client 123", true},
		{"too long client ID", strings.Repeat("a", 256), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateClientID(tt.clientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateClientID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMessage(t *testing.T) {
	tests := []struct {
		name    string
		message string
		wantErr bool
	}{
		{"valid message", "Hello World", false},
		{"empty message", "", true},
		{"too long message", strings.Repeat("a", 10001), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMessage(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNotification(t *testing.T) {
	validNotification := &notification.Notification{
		ClientID: "client-123",
		Message:  "Test message",
		Source:   "test",
	}

	tests := []struct {
		name         string
		notification *notification.Notification
		wantErr      bool
	}{
		{"valid notification", validNotification, false},
		{"nil notification", nil, true},
		{"invalid client ID", &notification.Notification{ClientID: "", Message: "test", Source: "test"}, true},
		{"invalid message", &notification.Notification{ClientID: "test", Message: "", Source: "test"}, true},
		{"invalid source", &notification.Notification{ClientID: "test", Message: "test", Source: ""}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNotification(tt.notification)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
