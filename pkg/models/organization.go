package models

import (
	"context"
	"time"
)

// Organization represents a git organization
type Organization struct {
	ID          string    `json:"-" db:"id"`
	Name        string    `json:"name" db:"name"`
	Avatar      string    `json:"avatar" db:"avatar"`
	GitProvider SCMDriver `json:"git_provider" db:"git_provider"`
	Created     time.Time `json:"-" db:"created_at"`
	Updated     time.Time `json:"-"  db:"updated_at"`
}

// OrganizationStore defines operations for working with organizations.
type OrganizationStore interface {
	// Find finds the organizations in the datastore by name and scm driver.
	Find(ctx context.Context, scmDriver SCMDriver, org *Organization) ([]*Organization, error)
}
