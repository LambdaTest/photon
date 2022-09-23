package errs

// ErrInvalidDriver is returned when SCM driver is not defined
var ErrInvalidDriver = New("Invalid Git SCM driver")

// ErrInvalidLoggerInstance is returned when logger instance is not supported.
var ErrInvalidLoggerInstance = New("Invalid logger instance")
