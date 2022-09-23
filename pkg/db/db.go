package db

import (
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/jmoiron/sqlx"
)

// DB is a pool of zero or more underlying connections to
// the photon database.
type DB struct {
	conn   *sqlx.DB
	logger lumber.Logger
}

// Execute and executes a function. Any error that is returned from the function is returned
// from the Execute() method.
func (db *DB) Execute(fn func(conn *sqlx.DB) error) (err error) {
	err = fn(db.conn)
	return err
}

// ExecuteTransaction executes a function within the context of a read-write managed
// transaction. If no error is returned from the function then the
// transaction is committed. If an error is returned then the entire
// transaction is rolled back. Any error that is returned from the function
// or returned from the commit is returned from the ExecuteTransaction() method.
func (db *DB) ExecuteTransaction(fn func(tx *sqlx.Tx) error) (err error) {
	tx, err := db.conn.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			if rerr := tx.Rollback(); rerr != nil {
				db.logger.Errorf("error while performing rollback, %v", err)
			}
			db.logger.Errorf("panic while executing query: %+v\n", p)
			panic(p) // More often than not a panic is due to a programming error, and should be corrected
		} else if err != nil {
			// something went wrong, rollback
			if rerr := tx.Rollback(); rerr != nil {
				db.logger.Errorf("error while performing rollback, %v", err)
			}
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()
	err = fn(tx)
	return err
}

// Close closes the database connection.
func (db *DB) Close() error {
	return db.conn.Close()
}
