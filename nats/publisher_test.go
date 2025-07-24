package nats

import (
	"encoding/json"
	"testing"
	"time"

	notification "github.com/MyWeHub/notification-sdk"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestPublishNotification(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL, nats.Timeout(500*time.Millisecond))
	if err != nil {
		t.Skip("Skipping test as no NATS server is available")
	}
	defer nc.Close()

	publisher, err := NewPublisher(nats.DefaultURL, "test-notifications")
	if err != nil {
		t.Fatalf("Failed to create notification publisher: %v", err)
	}
	defer publisher.Close()

	subject := "test-notifications.test-client"
	ch := make(chan *notification.Notification, 1)
	subscription, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		var n notification.Notification
		if err := json.Unmarshal(msg.Data, &n); err != nil {
			t.Errorf("Failed to unmarshal notification: %v", err)
			return
		}
		ch <- &n
	})
	if err != nil {
		t.Fatalf("Failed to subscribe to NATS: %v", err)
	}
	defer subscription.Unsubscribe()

	if err := nc.Flush(); err != nil {
		t.Fatalf("Failed to flush connection: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	err = publisher.PublishNotification("test-client", "Test message", notification.TypeInfo, "system")
	if err != nil {
		t.Fatalf("Failed to publish notification: %v", err)
	}

	select {
	case n := <-ch:
		assert.Equal(t, "test-client", n.ClientID)
		assert.Equal(t, "Test message", n.Message)
		assert.Equal(t, notification.TypeInfo, n.Type)
		assert.Equal(t, "system", n.Source)
		assert.False(t, n.Read)
		assert.NotEmpty(t, n.ID)
		assert.False(t, n.CreatedAt.IsZero())
	case <-time.After(3 * time.Second):
		t.Fatal("Timed out waiting for notification")
	}
}

func TestPublishCustomNotification(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL, nats.Timeout(500*time.Millisecond))
	if err != nil {
		t.Skip("Skipping test as no NATS server is available")
	}
	defer nc.Close()

	publisher, err := NewPublisher(nats.DefaultURL, "test-notifications")
	if err != nil {
		t.Fatalf("Failed to create notification publisher: %v", err)
	}
	defer publisher.Close()

	subject := "test-notifications.test-client"
	ch := make(chan *notification.Notification, 1)
	subscription, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		var n notification.Notification
		if err := json.Unmarshal(msg.Data, &n); err != nil {
			t.Errorf("Failed to unmarshal notification: %v", err)
			return
		}
		ch <- &n
	})
	if err != nil {
		t.Fatalf("Failed to subscribe to NATS: %v", err)
	}
	defer subscription.Unsubscribe()

	if err := nc.Flush(); err != nil {
		t.Fatalf("Failed to flush connection: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	customNotification := &notification.Notification{
		Message:  "Custom test message",
		Type:     notification.TypeWarning,
		ClientID: "test-client",
		Source:   "custom-test",
	}

	err = publisher.PublishCustomNotification("test-client", customNotification)
	if err != nil {
		t.Fatalf("Failed to publish custom notification: %v", err)
	}

	select {
	case n := <-ch:
		assert.Equal(t, "test-client", n.ClientID)
		assert.Equal(t, "Custom test message", n.Message)
		assert.Equal(t, notification.TypeWarning, n.Type)
		assert.Equal(t, "custom-test", n.Source)
		assert.False(t, n.Read)
		assert.NotEmpty(t, n.ID)
		assert.False(t, n.CreatedAt.IsZero())
	case <-time.After(3 * time.Second):
		t.Fatal("Timed out waiting for notification")
	}
}

func TestPublishNotificationValidation(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL, nats.Timeout(500*time.Millisecond))
	if err != nil {
		t.Skip("Skipping test as no NATS server is available")
	}
	defer nc.Close()

	publisher, err := NewPublisher(nats.DefaultURL, "test-notifications")
	if err != nil {
		t.Fatalf("Failed to create notification publisher: %v", err)
	}
	defer publisher.Close()

	err = publisher.PublishNotification("", "Test message", notification.TypeInfo, "system")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "clientID cannot be empty")

	err = publisher.PublishNotification("test-client", "", notification.TypeInfo, "system")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "message cannot be empty")

	err = publisher.PublishNotification("test-client", "Test message", notification.TypeInfo, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "source cannot be empty")
}

func TestPublishCustomNotificationValidation(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL, nats.Timeout(500*time.Millisecond))
	if err != nil {
		t.Skip("Skipping test as no NATS server is available")
	}
	defer nc.Close()

	publisher, err := NewPublisher(nats.DefaultURL, "test-notifications")
	if err != nil {
		t.Fatalf("Failed to create notification publisher: %v", err)
	}
	defer publisher.Close()

	err = publisher.PublishCustomNotification("", &notification.Notification{
		Message: "Test message",
		Type:    notification.TypeInfo,
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "clientID cannot be empty")

	err = publisher.PublishCustomNotification("test-client", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "notification cannot be nil")
}
