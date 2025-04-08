package domain

import "errors"

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrInsufficientFunds   = errors.New("insufficient funds")
	ErrInvalidTransaction  = errors.New("invalid transaction")
	ErrTransactionNotFound = errors.New("transaction not found")
	ErrDuplicateApiKey     = errors.New("api key already exists")
	ErrInvalidAmount       = errors.New("amount must be greater than 0")
	ErrInvalidStatus       = errors.New("invalid status")
	ErrSameStatus          = errors.New("cannot change invoice status to the same status")
	ErrInvoiceNotFound     = errors.New("invoice not found")
	ErrUnauthorizedAccess  = errors.New("unauthorized access")
)
