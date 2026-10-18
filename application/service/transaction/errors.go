package transaction

import apperrors "Momentum/application/errors"

var (
	ErrInvalidInput = apperrors.New(apperrors.CodeInvalidInput, 400, "invalid transaction input")
	ErrNotFound     = apperrors.New(apperrors.CodeNotFound, 404, "transaction not found")
)
