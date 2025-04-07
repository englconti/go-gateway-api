package domain

import "errors"

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrInsufficientFunds   = errors.New("insufficient funds")
	ErrInvalidTransaction  = errors.New("invalid transaction")
	ErrTransactionNotFound = errors.New("transaction not found")
	ErrDuplicateApiKey     = errors.New("Api key already exists")
)
