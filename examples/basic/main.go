package main

import (
	"log"

	notification "github.com/ahyaghoubi/notification-sdk"
	"github.com/ahyaghoubi/notification-sdk/nats"
)

func main() {
	// Create a new NATS publisher
	publisher, err := nats.NewPublisher("nats://localhost:4222", "notifications")
	if err != nil {
		log.Fatal("Failed to create publisher:", err)
	}
	defer publisher.Close()

	// Publish a simple notification
	err = publisher.PublishNotification("client-123", "Hello World!", notification.TypeInfo, "example-app")
	if err != nil {
		log.Fatal("Failed to publish notification:", err)
	}

	log.Println("Notification published successfully!")

	// Publish a custom notification
	customNotification := &notification.Notification{
		Message: "This is a custom notification",
		Type:    notification.TypeWarning,
		Source:  "custom-service",
		UserID:  "user-456",
	}

	err = publisher.PublishCustomNotification("client-123", customNotification)
	if err != nil {
		log.Fatal("Failed to publish custom notification:", err)
	}

	log.Println("Custom notification published successfully!")
}
