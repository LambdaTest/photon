package errs

import "database/sql"

var (
	// GenericErrorMessage is generic error message returned to user.
	GenericErrorMessage = New("Unexpected error. Please try again later.")

	// ErrSignatureInvalid is returned when the webhook
	// signature is invalid or cannot be calculated.
	ErrSignatureInvalid = New("Invalid webhook signature")

	// ErrUnknownEvent is returned when the webhook event
	// is not recognized by the system.
	ErrUnknownEvent = New("Unknown webhook event")

	// ErrQuery is returned when database query fails
	ErrQuery = New("SQL query failed")

	// ErrEmptyCommit is returned when there is empty commit event
	ErrEmptyCommit = New("Empty commit")

	// ErrPingEvent is returned when database query fails
	ErrPingEvent = New("webhook ping event")

	// ErrNotSupported is returned when database query fails
	ErrNotSupported = New("Event not supported")

	// ErrInvalidSCMProvider is returned when invalid scm driver is provided
	ErrInvalidSCMProvider = New("Invalid git scm provider")

	// ErrMarshalJSON is returned when json marshal failed
	ErrMarshalJSON = New("JSON marshal failed")

	// ErrRepoNotFound is returned when given repo is not active in database.
	ErrRepoNotFound = New("Repository not active.")

	// ErrRowsNotFound is returned by Scan when QueryRow doesn't return a
	// row.
	ErrRowsNotFound = sql.ErrNoRows
)
