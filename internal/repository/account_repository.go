package repository

import (
	"database/sql" // pacote do go para praticamente comunicar com qlq banco de dados
	"log"

	"time"

	"github.com/englconti/imersaoFC/go-gateway/internal/domain"
)

type AccountRepository struct {
	db *sql.DB // db é o sql.DB, que é o objeto que representa o banco de dados e a conexão com ele
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare(`
	INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	) // o exec é o comando para executar o sql, ele retorna um resultado que é a quantidade de linhas afetadas e um erro. Como não precisamos do resultado, usamos o _ para ignorar.
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	stmt, err := r.db.Prepare(`
	SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = $1
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var account domain.Account
	err = stmt.QueryRow(apiKey).Scan(
		&account.ID, // o & aqui serve para alterar o valor da variável account.ID, que é um ponteiro para o valor da variável account.ID
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		// log an error
		log.Printf("Account not found for API key: %s", apiKey)
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	stmt, err := r.db.Prepare(`
	SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE id = $1
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var account domain.Account
	err = stmt.QueryRow(id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // se der erro, o rollback é feito para desfazer as alterações

	var currentBalance float64
	err = tx.QueryRow(`
	SELECT balance FROM accounts WHERE id = $1 FOR UPDATE
	`, account.ID).Scan(&currentBalance) // FOR UPDATE é uma cláusula do sql que garante que a linha será bloqueada até que a transação seja concluída
	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3`, account.Balance, time.Now(), account.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
