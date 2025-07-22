package notification

import (
	"testing"
	"time"
)

func TestIsValid(t *testing.T) {
	// Create a valid notification
	notification := Notification{
		ID:        "1",
		ClientID:  "client1",
		UserID:    "user1",
		Message:   "Test message",
		Type:      TypeInfo,
		Read:      false,
		CreatedAt: time.Now(),
		Source:    "system",
	}

	// Test valid notification
	if !notification.IsValid() {
		t.Error("Expected notification to be valid")
	}

	// Test invalid notification (empty user ID)
	notification.UserID = ""
	if notification.IsValid() {
		t.Error("Expected notification to be invalid with empty user ID")
	}

	// Test invalid notification (empty message)
	notification.UserID = "user1"
	notification.Message = ""
	if notification.IsValid() {
		t.Error("Expected notification to be invalid with empty message")
	}

	// Test invalid notification (empty source)
	notification.Message = "Test message"
	notification.Source = ""
	if notification.IsValid() {
		t.Error("Expected notification to be invalid with empty source")
	}
}

func TestMarkAsRead(t *testing.T) {
	// Create a notification
	notification := Notification{
		ID:       "1",
		ClientID: "client1",
		UserID:   "user1",
		Message:  "Test message",
		Type:     TypeInfo,
		Read:     false,
	}

	// Mark the notification as read
	notification.MarkAsRead()

	// Verify the notification is marked as read
	if !notification.Read {
		t.Error("Expected notification to be marked as read")
	}
}
