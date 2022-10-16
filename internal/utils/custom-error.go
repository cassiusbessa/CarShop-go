package utils

import "fmt"

// Error defines a standard application error.
type CustomError struct {
	Code    int    // Machine-readable error code
	Message string // Human-readable message (final user)
	Op      string // Logical operation (operation role)
	Err     error  // Embedded error  (operation role)
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

// NewError creates a new custom error
func NewError(code int, message, op string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Op:      op,
		Err:     err,
	}
}
