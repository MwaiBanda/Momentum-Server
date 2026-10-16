package sermon

import "errors"

var (
	ErrInvalidPage = errors.New("invalid sermon page")
	ErrUnavailable = errors.New("sermon provider unavailable")
)
