package service

import (
	"github.com/englconti/imersaoFC/go-gateway/internal/domain"
	"github.com/englconti/imersaoFC/go-gateway/internal/dto"
)

type InvoiceService struct {
	repository     domain.InvoiceRepository
	accountService AccountService
}

func NewInvoiceService(repository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		repository:     repository,
		accountService: accountService,
	}
}

func (is *InvoiceService) Create(input dto.CreateInvoiceInput) (*dto.InvoiceOutput, error) {
	account, err := is.accountService.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	// Converte o input do request de JSON para uma struct Invoice para ser trabalhada pelo domínio
	invoice, err := dto.ToInvoice(input, account.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = is.accountService.UpdateBalance(account.APIKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err := is.repository.Save(invoice); err != nil {
		return nil, err
	}

	output := dto.FromInvoice(invoice)
	return &output, nil
}

func (is *InvoiceService) GetByID(id string, apiKey string) (*dto.InvoiceOutput, error) {
	invoice, err := is.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	account, err := is.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	if account.ID != invoice.AccountID {
		return nil, domain.ErrUnauthorizedAccess
	}

	output := dto.FromInvoice(invoice)
	return &output, nil
}

func (s *InvoiceService) ListByAccount(accountID string) ([]*dto.InvoiceOutput, error) {
	invoices, err := s.repository.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		invoiceOutput := dto.FromInvoice(invoice)
		output[i] = &invoiceOutput
	}

	return output, nil
}

// ListByAccountAPIKey
func (s *InvoiceService) ListByAccountAPIKey(apiKey string) ([]*dto.InvoiceOutput, error) {
	account, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccount(account.ID)
}
