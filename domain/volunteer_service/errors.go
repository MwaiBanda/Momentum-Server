package volunteer_service

import "errors"

var (
	ErrServiceNotFound = errors.New("service not found")
	ErrDayNotFound     = errors.New("volunteer service day not found")
	ErrInvalidInput    = errors.New("invalid volunteer service input")
)
