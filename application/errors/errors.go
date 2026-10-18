package errors

import (
	"errors"
	"fmt"
)

const (
	CodeInvalidInput          = "invalid_input"
	CodeInvalidRequest        = "invalid_request"
	CodeNotFound              = "not_found"
	CodeConflict              = "conflict"
	CodeUnauthorized          = "unauthorized"
	CodeForbidden             = "forbidden"
	CodeDependencyUnavailable = "dependency_unavailable"
	CodeInternal              = "internal_error"
)

type Error struct {
	code    string
	status  int
	message string
	cause   error
}

func New(code string, status int, message string) *Error {
	return &Error{
		code:    code,
		status:  status,
		message: message,
	}
}

func Wrap(cause error, code string, status int, message string) *Error {
	return &Error{
		code:    code,
		status:  status,
		message: message,
		cause:   cause,
	}
}

func InvalidInput(message string) *Error {
	return New(CodeInvalidInput, 400, message)
}

func (err *Error) Error() string {
	if err == nil {
		return ""
	}
	if err.cause != nil {
		return fmt.Sprintf("%s: %v", err.message, err.cause)
	}
	return err.message
}

func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.cause
}

func Details(err error) (code string, status int, message string, ok bool) {
	var applicationError *Error
	if !errors.As(err, &applicationError) || applicationError == nil {
		return "", 0, "", false
	}
	return applicationError.code, applicationError.status, applicationError.message, true
}
