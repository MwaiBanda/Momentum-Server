package event

import apperrors "Momentum/application/errors"

var ErrUnavailable = apperrors.New(
	apperrors.CodeDependencyUnavailable,
	503,
	"event provider unavailable",
)
