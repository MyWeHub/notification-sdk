package natsutil

import (
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConnectOptions(t *testing.T) {
	opts := DefaultConnectOptions()
	assert.NotEmpty(t, opts)
	assert.Len(t, opts, 5) // We have 5 default options
}

func TestConnectWithRetry(t *testing.T) {
	// Test with invalid URL - should fail after retries
	nc, err := ConnectWithRetry("nats://invalid-url:4222", 2)
	assert.Error(t, err)
	assert.Nil(t, nc)
	assert.Contains(t, err.Error(), "failed to connect to NATS after retries")
}

func TestConnectWithCustomOptions(t *testing.T) {
	// Test with custom timeout
	customOpts := []nats.Option{
		nats.Timeout(500 * time.Millisecond),
	}

	nc, err := ConnectWithCustomOptions("nats://invalid-url:4222", customOpts...)
	assert.Error(t, err)
	assert.Nil(t, nc)
}
