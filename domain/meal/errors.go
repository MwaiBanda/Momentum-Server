package meal

import "errors"

var (
	ErrNotFound                = errors.New("meal not found")
	ErrVolunteeredMealNotFound = errors.New("volunteered meal not found")
	ErrInvalidInput            = errors.New("invalid meal input")
)
