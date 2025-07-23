package main

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"

	notification "github.com/ahyaghoubi/notification-sdk"
	"github.com/ahyaghoubi/notification-sdk/nats"
)

func main() {
	// Initialize Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	// Create a new NATS publisher
	publisher, err := nats.NewPublisher("nats://localhost:4222", "notifications")
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal("Failed to create publisher:", err)
	}
	defer publisher.Close()

	// Publish a simple notification
	err = publisher.PublishNotification("client-123", "Hello World!", notification.TypeInfo, "example-app")
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal("Failed to publish notification:", err)
	}

	log.Println("Notification published successfully!")

	// Publish a custom notification
	customNotification := &notification.Notification{
		Message:  "This is a custom notification",
		Type:     notification.TypeWarning,
		Source:   "custom-service",
		UserID:   "user-456",
		ClientID: "client-123",
	}

	err = publisher.PublishCustomNotification("client-123", customNotification)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal("Failed to publish custom notification:", err)
	}

	log.Println("Custom notification published successfully!")
}
