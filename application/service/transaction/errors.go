package transaction

import "errors"

var (
	ErrInvalidInput = errors.New("invalid transaction input")
	ErrNotFound     = errors.New("transaction not found")
)
