package service

import (
	"github.com/englconti/imersaoFC/go-gateway/internal/domain"
	"github.com/englconti/imersaoFC/go-gateway/internal/dto"
)

type AccountService struct { // um struct é uma coleção de campos e métodos. Ele serve para agrupar dados e comportamentos relacionados. A diferença entre struct e interface é que struct é uma implementação concreta de uma interface.
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (as *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := as.repository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrDuplicateApiKey
	}

	err = as.repository.Save(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)

	return &output, nil
}

func (as *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := as.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = as.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (as *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := as.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account) // output é necessário para retornar o valor de AccountOutput porque o FromAccount retorna um AccountOutput e não um ponteiro para AccountOutput
	return &output, nil
}

func (as *AccountService) FindByID(id string) (*dto.AccountOutput, error) {
	account, err := as.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}
