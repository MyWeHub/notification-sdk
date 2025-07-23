package utils

import (
	"time"
)

// UTCNow returns the current time in UTC
func UTCNow() time.Time {
	return time.Now().UTC()
}

// IsZeroTime checks if a time is zero
func IsZeroTime(t time.Time) bool {
	return t.IsZero()
}

// FormatTimestamp formats a time as RFC3339
func FormatTimestamp(t time.Time) string {
	return t.Format(time.RFC3339)
}
