package natsutil

import (
	"fmt"
	"strings"
)

// BuildSubject constructs a NATS subject from prefix and client ID
func BuildSubject(prefix, clientID string) string {
	// Sanitize the client ID for use in NATS subjects
	sanitized := SanitizeForSubject(clientID)
	return fmt.Sprintf("%s.%s", prefix, sanitized)
}

// SanitizeForSubject removes invalid characters from a string for NATS subject use
func SanitizeForSubject(input string) string {
	// Replace spaces and special characters with underscores
	result := strings.ReplaceAll(input, " ", "_")
	result = strings.ReplaceAll(result, ".", "_")
	result = strings.ReplaceAll(result, "*", "_")
	result = strings.ReplaceAll(result, ">", "_")
	return result
}

// ValidateSubject checks if a subject is valid for NATS
func ValidateSubject(subject string) bool {
	if subject == "" {
		return false
	}

	// NATS subjects cannot contain spaces
	if strings.Contains(subject, " ") {
		return false
	}

	return true
}
