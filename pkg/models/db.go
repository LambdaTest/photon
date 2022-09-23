package models

import "github.com/jmoiron/sqlx"

// DB represents the MySQL client
type DB interface {
	// Close closes the db connection.
	Close() error
	// ExecuteTransaction wrapper for executing queries in a transaction.
	ExecuteTransaction(fn func(tx *sqlx.Tx) error) (err error)
	// Execute wrapper for executing the queries.
	Execute(fn func(conn *sqlx.DB) error) (err error)
}
