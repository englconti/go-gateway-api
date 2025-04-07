package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/englconti/imersaoFC/go-gateway/internal/repository"
	"github.com/englconti/imersaoFC/go-gateway/internal/service"
	"github.com/englconti/imersaoFC/go-gateway/internal/web/server"
	_ "github.com/lib/pq" // driver do postgres. O _ serve para que o compilador não reclame que o driver não é usado.

	"github.com/joho/godotenv"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// string de conexão com o banco de dados
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	defer db.Close()
	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("PORT", "8080")
	srv := server.NewServer(port, accountService)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
