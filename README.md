# Notification SDK

A modular, hexagonal-architecture Go SDK for publishing notifications via NATS.

## Features
- Hexagonal (ports & adapters) architecture
- Core domain types for notifications
- NATS adapter for publishing notifications
- Easily extensible for other transports

## Project Structure
```
notification-sdk/
  domain/                # Core domain types (Notification, errors)
  port/                  # Port interfaces (NotificationPublisherPort)
  adapter/nats/          # NATS implementation of the publisher
  test/                  # Tests for the SDK
  go.mod, go.sum         # Go module files
```

## Usage

### 1. Install
```sh
go get github.com/ahyaghoubi/notification-sdk
```

### 2. Import and Use
```go
import (
    "github.com/ahyaghoubi/notification-sdk/adapter/nats"
    "github.com/ahyaghoubi/notification-sdk/domain"
)

func main() {
    publisher, err := nats.NewNotificationPublisher(nats.DefaultURL, "notifications")
    if err != nil {
        log.Fatal(err)
    }
    defer publisher.Close()

    err = publisher.PublishNotification("client-id", "Hello!", domain.TypeInfo, "system")
    if err != nil {
        log.Fatal(err)
    }
}
```

### 3. Run Tests
```sh
go test ./test/...
```

## Extending
- Implement your own adapter by creating a new package in `adapter/` and implementing the `port.NotificationPublisherPort` interface.

## License
MIT 