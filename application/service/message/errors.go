package message

import apperrors "Momentum/application/errors"

var (
	ErrMessageNotFound = apperrors.New(apperrors.CodeNotFound, 404, "message not found")
	ErrPassageNotFound = apperrors.New(apperrors.CodeNotFound, 404, "passage not found")
	ErrNoteNotFound    = apperrors.New(apperrors.CodeNotFound, 404, "note not found")
	ErrInvalidInput    = apperrors.New(apperrors.CodeInvalidInput, 400, "invalid message input")
)
