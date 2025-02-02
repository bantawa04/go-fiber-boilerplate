package errors

import (
    "errors"
    "fmt"
)

// Custom error types
type AppError struct {
    Code    int
    Message string
    Err     error
}

func (e *AppError) Error() string {
    return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

// Error constructors
func NewNotFoundError(message string) *AppError {
    return &AppError{
        Code:    404,
        Message: message,
        Err:     errors.New("not found"),
    }
}

func NewBadRequestError(message string) *AppError {
    return &AppError{
        Code:    400,
        Message: message,
        Err:     errors.New("bad request"),
    }
}

func NewConflictError(message string) *AppError {
    return &AppError{
        Code:    409,
        Message: message,
        Err:     errors.New("conflict"),
    }
}

func NewInternalError(err error) *AppError {
    return &AppError{
        Code:    500,
        Message: "Internal server error",
        Err:     err,
    }
}