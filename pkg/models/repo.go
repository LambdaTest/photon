package models

import (
	"context"
	"time"
)

// Repository represents a git repository
type Repository struct {
	ID                string    `json:"id,omitempty" db:"id"`
	OrgID             string    `json:"-" db:"org_id"`
	Strict            bool      `json:"-" db:"strict"`
	Name              string    `json:"name,omitempty" db:"name"`
	Admin             string    `json:"-" db:"admin_id"`
	Namespace         string    `json:"namespace,omitempty"`
	Private           bool      `json:"-"`
	Secret            string    `json:"-" db:"webhook_secret"`
	Link              string    `json:"link" db:"link"`
	HTTPURL           string    `json:"http_url,omitempty" db:"git_http_url"`
	SSHURL            string    `json:"ssh_url,omitempty" db:"git_ssh_url"`
	Active            bool      `json:"active" db:"active"`
	TasFileName       string    `json:"-"  db:"tas_file_name"`
	PostMergeStrategy int       `json:"-"  db:"post_merge_strategy"`
	Created           time.Time `json:"-"  db:"created_at"`
	Updated           time.Time `json:"-"  db:"updated_at"`
	Perm              *Perm     `json:"permissions,omitempty"`
	Mask              string    `json:"-" db:"mask"`
}

// Perm represents the user' s repository permissions
type Perm struct {
	Read  bool `db:"perm_read"     json:"read"`
	Write bool `db:"perm_write"    json:"write"`
	Admin bool `db:"perm_admin"    json:"admin"`
}

// RepoStore defines operations for working with repositories.
type RepoStore interface {
	// FindActiveByName returns the active repositories in the data store by commitID.
	FindActiveByName(ctx context.Context, repoName, orgName string, gitProvider SCMDriver) (repo *Repository, err error)
}
