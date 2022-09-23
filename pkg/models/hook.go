package models

import (
	"net/http"

	"github.com/segmentio/kafka-go"

	"github.com/drone/go-scm/scm"
)

// EventType represents the webhook event
type EventType string

const (
	// EventPush represents the push event.
	EventPush EventType = "push"
	// EventPullRequest represents the pull request event.
	EventPullRequest EventType = "pull-request"
	// EventPing represents the ping event.
	EventPing EventType = "ping"
)

// HookParser parses the webhook from the source
// code management system.
type HookParser interface {
	Parse(req *http.Request, driver SCMDriver) (scm.Webhook, []kafka.Header, error)
}
