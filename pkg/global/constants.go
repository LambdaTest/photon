package global

import "time"

const (
	// ServiceName OpenTelemetry service name
	ServiceName = "photon"
	// MysqlMaxIdleConnection max mysql idle connections.
	MysqlMaxIdleConnection = 25
	// MysqlMaxOpenConnection max mysql open connections.
	MysqlMaxOpenConnection = 25
	// MysqlMaxConnectionLifetime max mysql connection lifetime.
	MysqlMaxConnectionLifetime = 5 * time.Minute
)
