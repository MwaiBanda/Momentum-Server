package notification

import apperrors "Momentum/application/errors"

var ErrUnavailable = apperrors.New(
	apperrors.CodeDependencyUnavailable,
	503,
	"notification sender unavailable",
)
