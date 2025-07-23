# Notification SDK

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go) 
![License](https://img.shields.io/badge/License-MIT-blue.svg)
![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)

A modular, hexagonal-architecture Go SDK for publishing notifications via NATS and other message brokers.

## 🚀 Features

- **Clean Architecture** - Follows hexagonal (ports & adapters) design principles
- **Type-Safe** - Full Go type safety with structured error handling
- **NATS Integration** - Production-ready NATS publisher with JetStream support
- **Extensible** - Easy to add support for other message brokers (Kafka, RabbitMQ, etc.)
- **Validation** - Built-in validation for all notification fields
- **Auto-Fill** - Automatic generation of IDs, timestamps, and other metadata
- **Organization Support** - Built-in support for organization preferences
- **Testing Ready** - Comprehensive test suite and integration test helpers

## 📦 Installation

```
go get github.com/ahyaghoubi/notification-sdk@latest
```

## 🏃 Quick Start

```
package main

import (
    "log"
    
    notification "github.com/ahyaghoubi/notification-sdk"
    "github.com/ahyaghoubi/notification-sdk/nats"
)

func main() {
    // Create NATS publisher
    publisher, err := nats.NewPublisher("nats://localhost:4222", "notifications")
    if err != nil {
        log.Fatal("Failed to create publisher:", err)
    }
    defer publisher.Close()

    // Publish a notification
    err = publisher.PublishNotification(
        "client-123",              // Client ID
        "Hello World!",            // Message
        notification.TypeInfo,     // Notification type
        "my-service",             // Source service
    )
    if err != nil {
        log.Fatal("Failed to publish:", err)
    }

    log.Println("✅ Notification published successfully!")
}
```

## 📖 Usage Guide

### Basic Notification Publishing

```
publisher, err := nats.NewPublisher("nats://localhost:4222", "notifications")
if err != nil {
    return err
}
defer publisher.Close()

// Simple notification
err = publisher.PublishNotification(
    "user-456",                    // Client/User ID
    "Your order has been shipped", // Message
    notification.TypeSuccess,      // Type
    "order-service",              // Source
)
if err != nil {
    log.Printf("Failed to publish: %v", err)
}
```

### Custom Notifications

```
// Create a custom notification with additional fields
customNotification := &notification.Notification{
    Message: "System maintenance scheduled",
    Type:    notification.TypeWarning,
    Source:  "admin-panel",
    UserID:  "admin-001",
    // ID and CreatedAt will be auto-generated if not provided
}

err = publisher.PublishCustomNotification("broadcast", customNotification)
if err != nil {
    log.Printf("Failed to publish custom notification: %v", err)
}
```

### Advanced NATS Configuration

```
import "github.com/nats-io/nats.go"

// Custom NATS options
customOpts := []nats.Option{
    nats.Token("your-auth-token"),
    nats.Timeout(30 * time.Second),
    nats.MaxReconnects(5),
}

publisher, err := nats.NewPublisherWithOptions(
    "nats://auth-server:4222", 
    "notifications",
    customOpts...,
)
```

### Error Handling

```
err := publisher.PublishNotification("", "message", notification.TypeInfo, "source")
if err != nil {
    // Type assertion to get detailed error information
    if notifErr, ok := err.(*notification.Error); ok {
        switch notifErr.Code {
        case notification.InvalidArguments:
            log.Printf("Validation error: %s", notifErr.Message)
        case notification.Internal:
            log.Printf("Internal error: %s", notifErr.Message)
        default:
            log.Printf("Unknown error: %s", notifErr.Message)
        }
    }
}
```

## 🏗️ Architecture

This SDK follows **Hexagonal Architecture** principles:

```
notification-sdk/
├── types.go              # 🏛️  Core domain types and errors
├── interfaces.go         # 🔌  Port definitions (interfaces)
├── nats/                 # 🔄  NATS adapter implementation
│   ├── publisher.go
│   └── publisher_test.go
├── internal/             # 🔒  Private utilities (not importable)
│   ├── validation/       # ✅  Input validation logic
│   ├── utils/           # 🛠️  JSON, time utilities
│   └── natsutil/        # 📡  NATS connection helpers
└── examples/            # 📚  Usage examples
    └── basic/main.go
```

### Core Concepts

- **Domain Types** (`types.go`) - Business entities like `Notification`, `NotificationType`
- **Ports** (`interfaces.go`) - Abstract interfaces like `PublisherPort`
- **Adapters** (`nats/`, `kafka/`) - Concrete implementations for different technologies
- **Internal** (`internal/`) - Private utilities hidden from consumers

## 📋 API Reference

### Notification Types

```
const (
    TypeInfo    NotificationType = 0  // ℹ️  Informational messages
    TypeWarning NotificationType = 1  // ⚠️  Warning messages
    TypeError   NotificationType = 2  // ❌ Error messages
    TypeSuccess NotificationType = 3  // ✅ Success messages
    TypeSystem  NotificationType = 4  // 🔧 System messages
)
```

### Notification Structure

```
type Notification struct {
    ID        string           `json:"id"`         // Auto-generated UUID
    ClientID  string           `json:"client_id"`  // Target client/user
    UserID    string           `json:"user_id"`    // Optional user ID
    Message   string           `json:"message"`    // Notification content
    Type      NotificationType `json:"type"`       // Notification type
    Read      bool             `json:"read"`       // Read status
    CreatedAt time.Time        `json:"created_at"` // Auto-generated timestamp
    Source    string           `json:"source"`     // Source service name
}
```

### Publisher Interface

```
type PublisherPort interface {
    // Publish a simple notification
    PublishNotification(clientID, message string, notificationType NotificationType, source string) error
    
    // Publish a custom notification with full control
    PublishCustomNotification(clientID string, notification *Notification) error
    
    // Close the publisher and cleanup resources
    Close() error
}
```

### Error Codes

```
const (
    NotFound         = 404  // Resource not found
    InvalidArguments = 400  // Validation errors
    AlreadyExists    = 409  // Duplicate resource
    PermissionDenied = 403  // Access denied
    Unauthorized     = 401  // Authentication required
    Internal         = 500  // Internal system error
)
```

## 🧪 Testing

### Run All Tests

```
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests verbosely
go test -v ./...
```

### Integration Tests

The NATS tests require a running NATS server:

```
# Install NATS server
go install github.com/nats-io/nats-server/v2@latest

# Run NATS server
nats-server

# Run integration tests
go test ./nats -v
```

### Using Docker for Testing

```
# Start NATS server with Docker
docker run -d --name nats -p 4222:4222 nats:latest

# Run tests
go test ./...

# Cleanup
docker stop nats && docker rm nats
```

## 🔧 Examples

### Basic Example

```
// examples/basic/main.go
package main

import (
    "log"
    notification "github.com/ahyaghoubi/notification-sdk"
    "github.com/ahyaghoubi/notification-sdk/nats"
)

func main() {
    publisher, err := nats.NewPublisher("nats://localhost:4222", "notifications")
    if err != nil {
        log.Fatal(err)
    }
    defer publisher.Close()

    // Simple notification
    err = publisher.PublishNotification("user-123", "Welcome!", notification.TypeInfo, "auth-service")
    if err != nil {
        log.Fatal(err)
    }

    // Custom notification
    custom := &notification.Notification{
        Message: "Your profile was updated",
        Type:    notification.TypeSuccess,
        Source:  "profile-service",
        UserID:  "user-123",
    }
    
    err = publisher.PublishCustomNotification("user-123", custom)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("✅ All notifications sent successfully!")
}
```

### Batch Publishing Example

```
func publishBatch(publisher notification.PublisherPort, notifications []NotificationData) error {
    for _, data := range notifications {
        err := publisher.PublishNotification(
            data.ClientID,
            data.Message,
            data.Type,
            data.Source,
        )
        if err != nil {
            return fmt.Errorf("failed to publish to %s: %w", data.ClientID, err)
        }
    }
    return nil
}
```

### Error Handling Example

```
func handlePublishError(err error) {
    if err == nil {
        return
    }

    if notifErr, ok := err.(*notification.Error); ok {
        switch notifErr.Code {
        case notification.InvalidArguments:
            log.Printf("❌ Validation failed: %s", notifErr.Message)
            // Maybe retry with corrected data
        case notification.Internal:
            log.Printf("🔧 System error: %s", notifErr.Message)
            // Maybe implement retry logic
        case notification.Unauthorized:
            log.Printf("🔐 Auth error: %s", notifErr.Message)
            // Refresh credentials
        default:
            log.Printf("⚠️  Unknown error: %s", notifErr.Message)
        }
    } else {
        log.Printf("🚨 Unexpected error: %v", err)
    }
}
```

## 🔌 Extending the SDK

### Adding a New Adapter (e.g., Kafka)

1. **Create the adapter package:**

```
mkdir kafka
```

2. **Implement the PublisherPort interface:**

```
// kafka/publisher.go
package kafka

import (
    notification "github.com/ahyaghoubi/notification-sdk"
)

type Publisher struct {
    // Kafka-specific fields
}

func NewPublisher(brokers []string, topic string) (*Publisher, error) {
    // Kafka connection logic
    return &Publisher{}, nil
}

func (p *Publisher) PublishNotification(clientID, message string, notificationType notification.NotificationType, source string) error {
    // Kafka publishing logic
    return nil
}

func (p *Publisher) PublishCustomNotification(clientID string, notif *notification.Notification) error {
    // Kafka custom publishing logic
    return nil
}

func (p *Publisher) Close() error {
    // Cleanup Kafka resources
    return nil
}
```

3. **Add tests:**

```
// kafka/publisher_test.go
package kafka

import (
    "testing"
    notification "github.com/ahyaghoubi/notification-sdk"
)

func TestKafkaPublisher(t *testing.T) {
    // Test implementation
}
```

4. **Usage:**

```
import "github.com/ahyaghoubi/notification-sdk/kafka"

publisher, err := kafka.NewPublisher([]string{"localhost:9092"}, "notifications")
// Use the same PublisherPort interface
```

### Adding Custom Validation

```
// internal/validation/custom.go
package validation

func ValidateCustomField(field string) error {
    // Custom validation logic
    return nil
}
```

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork the repository**
2. **Create a feature branch:** `git checkout -b feature/amazing-feature`
3. **Make your changes** and add tests
4. **Run tests:** `go test ./...`
5. **Commit your changes:** `git commit -m 'Add amazing feature'`
6. **Push to branch:** `git push origin feature/amazing-feature`
7. **Open a Pull Request**

### Development Setup

```
# Clone the repository
git clone https://github.com/ahyaghoubi/notification-sdk.git
cd notification-sdk

# Install dependencies
go mod download

# Run tests
go test ./...

# Run linter (if available)
golangci-lint run
```

