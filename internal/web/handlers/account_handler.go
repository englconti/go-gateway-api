package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/englconti/imersaoFC/go-gateway/internal/dto"
	"github.com/englconti/imersaoFC/go-gateway/internal/service"
)

type AccountHandler struct {
	service *service.AccountService
}

// NewAccountHandler cria um novo AccountHandler. É a função construtora da struct AccountHandler. Serve para criar um novo AccountHandler.
func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

// Abaixo a interface para receber e tratar as requisições HTTP.

func (ah *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	// w é o response writer e r é a requisição. O response writer é o objeto que escreve a resposta HTTP.
	var input dto.CreateAccountInput
	err := json.NewDecoder(r.Body).Decode(&input) // Decode decodifica o corpo da requisição em um objeto.
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := ah.service.CreateAccount(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// essa parte abaixo é para escrever a resposta HTTP.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (ah *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "X-API-Key header is required", http.StatusBadRequest)
		return
	}

	output, err := ah.service.FindByAPIKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // StatusOK é o status code para uma requisição bem sucedida. Entretanto ela já é padrão, então não seria necessário escrever.
	json.NewEncoder(w).Encode(output)
}
