package volunteer_service

import apperrors "Momentum/application/errors"

var (
	ErrServiceNotFound = apperrors.New(apperrors.CodeNotFound, 404, "service not found")
	ErrDayNotFound     = apperrors.New(apperrors.CodeNotFound, 404, "volunteer service day not found")
	ErrInvalidInput    = apperrors.New(apperrors.CodeInvalidInput, 400, "invalid volunteer service input")
)
