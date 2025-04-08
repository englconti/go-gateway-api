package domain

// Esse arquivo define as interfaces de repositório para as entidades Account e Invoice.
// Essas interfaces são usadas para abstrair o acesso ao banco de dados e facilitar a substituição de implementações.
type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	FindByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}

type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindByID(id string) (*Invoice, error)
	FindByAccountID(accountID string) ([]*Invoice, error) // []*Invoice retorna um slice(lista) de invoices
	UpdateStatus(invoice *Invoice) error
}
