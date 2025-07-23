package natsutil

import (
	"time"

	notification "github.com/ahyaghoubi/notification-sdk"
	"github.com/nats-io/nats.go"
)

// DefaultConnectOptions returns default NATS connection options
func DefaultConnectOptions() []nats.Option {
	return []nats.Option{
		nats.MaxReconnects(-1),                                    // Unlimited reconnections
		nats.ReconnectWait(2 * time.Second),                       // Wait 2 seconds between reconnect attempts
		nats.ReconnectJitter(500*time.Millisecond, 2*time.Second), // Add jitter to reconnect attempts
		nats.ReconnectBufSize(8 * 1024 * 1024),                    // 8MB reconnect buffer
		nats.Timeout(10 * time.Second),                            // Connection timeout
	}
}

// ConnectWithRetry connects to NATS with retry logic
func ConnectWithRetry(url string, maxRetries int) (*nats.Conn, error) {
	opts := DefaultConnectOptions()

	var lastErr error
	for i := 0; i < maxRetries; i++ {
		nc, err := nats.Connect(url, opts...)
		if err == nil {
			return nc, nil
		}

		lastErr = err
		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}
	}

	return nil, notification.NewError(notification.Internal, "failed to connect to NATS after retries: "+lastErr.Error())
}

// ConnectWithCustomOptions connects to NATS with custom options
func ConnectWithCustomOptions(url string, opts ...nats.Option) (*nats.Conn, error) {
	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, notification.NewError(notification.Internal, "failed to connect to NATS: "+err.Error())
	}
	return nc, nil
}

// CreateJetStreamContext creates a JetStream context from a NATS connection
func CreateJetStreamContext(nc *nats.Conn) (nats.JetStreamContext, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, notification.NewError(notification.Internal, "failed to create JetStream context: "+err.Error())
	}
	return js, nil
}
