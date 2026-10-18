package sermon

import apperrors "Momentum/application/errors"

var (
	ErrInvalidPage = apperrors.New(apperrors.CodeInvalidInput, 400, "invalid sermon page")
	ErrUnavailable = apperrors.New(
		apperrors.CodeDependencyUnavailable,
		503,
		"sermon provider unavailable",
	)
)
