package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

var validStatusesToGoFromPending = map[Status]bool{
	StatusApproved: true,
	StatusRejected: true,
}

type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardHolderName string
}

func NewInvoice(accountID string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.Number[len(card.Number)-4:] // pega os ultimos 4 digitos do numero do cartão

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func generateRandomStatus() Status {
	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	if randomSource.Float64() > 0.7 {
		return StatusApproved
	}
	return StatusRejected
}

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	i.Status = generateRandomStatus()
	return nil
}

func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status == newStatus {
		return ErrSameStatus
	}
	if !validStatusesToGoFromPending[newStatus] {
		return ErrInvalidStatus
	}
	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
