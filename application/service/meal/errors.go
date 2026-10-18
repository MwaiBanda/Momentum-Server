package meal

import apperrors "Momentum/application/errors"

var (
	ErrNotFound                = apperrors.New(apperrors.CodeNotFound, 404, "meal not found")
	ErrVolunteeredMealNotFound = apperrors.New(apperrors.CodeNotFound, 404, "volunteered meal not found")
	ErrInvalidInput            = apperrors.New(apperrors.CodeInvalidInput, 400, "invalid meal input")
)
