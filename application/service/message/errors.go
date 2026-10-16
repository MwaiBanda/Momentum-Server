package message

import "errors"

var (
	ErrMessageNotFound = errors.New("message not found")
	ErrPassageNotFound = errors.New("passage not found")
	ErrNoteNotFound    = errors.New("note not found")
	ErrInvalidInput    = errors.New("invalid message input")
)
