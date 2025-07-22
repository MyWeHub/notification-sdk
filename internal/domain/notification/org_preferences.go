package notification

type WorkflowEmailPreference struct {
	Enabled bool `json:"enabled"`
}

type OrganizationNotificationPreferences struct {
	ID             string                             `json:"id"`
	OrgID          string                             `json:"org_id"`
	InternalEmails []string                           `json:"internal_emails"`
	ExternalEmails []string                           `json:"external_emails"`
	Workflows      map[string]WorkflowEmailPreference `json:"workflows"`
}
