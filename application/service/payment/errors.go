package payment

import apperrors "Momentum/application/errors"

var ErrInvalidInput = apperrors.New(
	apperrors.CodeInvalidInput,
	400,
	"invalid payment input",
)
