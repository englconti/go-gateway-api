package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid" // importa direto do namespace do pacote
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mu        sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateAPIKey() string {
	b := make([]byte, 16) // cria um slice de 16 bytes. Slice é uma estrutura de dados que representa um array de bytes ou um array que pode ser dinâmico, mudar de tamanho.
	rand.Read(b)          // preenche o slice com bytes aleatórios

	return hex.EncodeToString(b) // converte o slice de bytes para uma string hexadecimal
}

func NewAccount(name, email string) *Account { // *Account é um ponteiro para um Account e é o valor de retorno da função
	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		APIKey:    generateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (a *Account) AddBalance(amount float64) {
	a.mu.Lock()         // bloqueia o acesso ao balance
	defer a.mu.Unlock() // libera o acesso ao balance após a execução da função. O defer é usado para garantir que o acesso ao balance seja liberado após a execução da função.

	a.Balance += amount
	a.UpdatedAt = time.Now()
}
