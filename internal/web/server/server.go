package server

import (
	"net/http"

	"github.com/englconti/imersaoFC/go-gateway/internal/service"
	"github.com/englconti/imersaoFC/go-gateway/internal/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux // chi é um router para o Go. É uma biblioteca que cria um router para o Go.
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(port string, accountService *service.AccountService) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
