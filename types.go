package notification

import "time"

// Error codes
const (
	NotFound         = 404
	InvalidArguments = 400
	AlreadyExists    = 409
	PermissionDenied = 403
	Unauthorized     = 401
	Internal         = 500
)

// Error represents a domain error with a code and message
type Error struct {
	Code    int32
	Message string
}

// NewError creates a new Error
func NewError(code int32, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Error implements the error interface
func (e Error) Error() string {
	return e.Message
}

// NotificationType defines the type of notification
type NotificationType int32

// Notification types
const (
	TypeInfo    NotificationType = 0
	TypeWarning NotificationType = 1
	TypeError   NotificationType = 2
	TypeSuccess NotificationType = 3
	TypeSystem  NotificationType = 4
)

// Notification represents a message sent to a user
type Notification struct {
	ID        string           `json:"id"`
	ClientID  string           `json:"client_id"`
	UserID    string           `json:"user_id"`
	Message   string           `json:"message"`
	Type      NotificationType `json:"type"`
	Read      bool             `json:"read"`
	CreatedAt time.Time        `json:"created_at"`
	Source    string           `json:"source"`
}

// NotificationEvent represents a notification with an event ID for SSE
type NotificationEvent struct {
	Notification *Notification
	EventID      string
}

// MarkAsRead marks the notification as read
func (n *Notification) MarkAsRead() {
	n.Read = true
}

// IsValid checks if the notification has the required fields
func (n *Notification) IsValid() bool {
	return n.UserID != "" && n.Message != "" && n.Source != ""
}

// WorkflowEmailPreference represents email preferences for a workflow
type WorkflowEmailPreference struct {
	Enabled bool `json:"enabled"`
}

// OrganizationNotificationPreferences represents notification preferences for an organization
type OrganizationNotificationPreferences struct {
	ID             string                             `json:"id"`
	OrgID          string                             `json:"org_id"`
	InternalEmails []string                           `json:"internal_emails"`
	ExternalEmails []string                           `json:"external_emails"`
	Workflows      map[string]WorkflowEmailPreference `json:"workflows"`
}
