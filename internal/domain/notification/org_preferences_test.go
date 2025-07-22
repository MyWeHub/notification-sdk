package notification

import (
	"encoding/json"
	"testing"
)

func TestOrganizationNotificationPreferences_Basic(t *testing.T) {
	prefs := OrganizationNotificationPreferences{
		OrgID:          "org1",
		InternalEmails: []string{"internal@example.com"},
		ExternalEmails: []string{"external@example.com"},
	}
	if prefs.OrgID != "org1" {
		t.Errorf("expected OrgID 'org1', got '%s'", prefs.OrgID)
	}
	if len(prefs.InternalEmails) != 1 {
		t.Errorf("expected 1 internal email, got %d", len(prefs.InternalEmails))
	}
	if len(prefs.ExternalEmails) != 1 {
		t.Errorf("expected 1 external email, got %d", len(prefs.ExternalEmails))
	}
}

func TestOrganizationNotificationPreferences_EmptyEmails(t *testing.T) {
	prefs := OrganizationNotificationPreferences{
		OrgID:          "org2",
		InternalEmails: []string{},
		ExternalEmails: []string{},
	}
	if prefs.OrgID != "org2" {
		t.Errorf("expected OrgID 'org2', got '%s'", prefs.OrgID)
	}
	if len(prefs.InternalEmails) != 0 {
		t.Errorf("expected 0 internal emails, got %d", len(prefs.InternalEmails))
	}
	if len(prefs.ExternalEmails) != 0 {
		t.Errorf("expected 0 external emails, got %d", len(prefs.ExternalEmails))
	}
}

func TestOrganizationNotificationPreferences_JSON(t *testing.T) {
	prefs := OrganizationNotificationPreferences{
		OrgID:          "org3",
		InternalEmails: []string{"internal@example.com"},
		ExternalEmails: []string{"external@example.com"},
	}
	data, err := json.Marshal(prefs)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}
	var out OrganizationNotificationPreferences
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if out.OrgID != prefs.OrgID ||
		len(out.InternalEmails) != 1 || out.InternalEmails[0] != "internal@example.com" ||
		len(out.ExternalEmails) != 1 || out.ExternalEmails[0] != "external@example.com" {
		t.Errorf("unexpected round-trip result: %+v", out)
	}
}

func TestOrganizationNotificationPreferences_WorkflowEmailPreference_Default(t *testing.T) {
	prefs := OrganizationNotificationPreferences{
		OrgID:          "org1",
		InternalEmails: []string{"internal@example.com"},
		ExternalEmails: []string{"external@example.com"},
		Workflows:      nil,
	}
	_, exists := prefs.Workflows["wf-1"]
	if exists {
		t.Error("expected workflow not to exist, but it does")
	}
}

func TestOrganizationNotificationPreferences_WorkflowEmailPreference_EnableDisable(t *testing.T) {
	prefs := OrganizationNotificationPreferences{
		OrgID:          "org1",
		InternalEmails: []string{"internal@example.com"},
		ExternalEmails: []string{"external@example.com"},
		Workflows:      map[string]WorkflowEmailPreference{},
	}
	id := "wf-2"
	prefs.Workflows[id] = WorkflowEmailPreference{Enabled: true}
	if !prefs.Workflows[id].Enabled {
		t.Error("expected workflow email to be enabled")
	}
	prefs.Workflows[id] = WorkflowEmailPreference{Enabled: false}
	if prefs.Workflows[id].Enabled {
		t.Error("expected workflow email to be disabled")
	}
}

func TestOrganizationNotificationPreferences_WorkflowEmailPreference_JSON(t *testing.T) {
	prefs := OrganizationNotificationPreferences{
		OrgID:          "org3",
		InternalEmails: []string{"internal@example.com"},
		ExternalEmails: []string{"external@example.com"},
		Workflows:      map[string]WorkflowEmailPreference{"wf-3": {Enabled: true}},
	}
	data, err := json.Marshal(prefs)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}
	var out OrganizationNotificationPreferences
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if !out.Workflows["wf-3"].Enabled {
		t.Errorf("expected workflow wf-3 to be enabled after round-trip")
	}
}

func TestOrganizationNotificationPreferences_SeparateEmailTypes(t *testing.T) {
	prefs := OrganizationNotificationPreferences{
		OrgID:          "org4",
		InternalEmails: []string{"admin@company.com", "dev@company.com"},
		ExternalEmails: []string{"client@external.com", "partner@vendor.com"},
	}

	if len(prefs.InternalEmails) != 2 {
		t.Errorf("expected 2 internal emails, got %d", len(prefs.InternalEmails))
	}
	if len(prefs.ExternalEmails) != 2 {
		t.Errorf("expected 2 external emails, got %d", len(prefs.ExternalEmails))
	}

	// Test JSON marshaling with separate email types
	data, err := json.Marshal(prefs)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var out OrganizationNotificationPreferences
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(out.InternalEmails) != 2 || len(out.ExternalEmails) != 2 {
		t.Errorf("unexpected email counts after round-trip: internal=%d, external=%d",
			len(out.InternalEmails), len(out.ExternalEmails))
	}
}
