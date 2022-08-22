package main

import (
	"fmt"
	"github.com/EmirShimshir/crud-books/internal/config"
	"github.com/EmirShimshir/crud-books/internal/repository/psql"
	"github.com/EmirShimshir/crud-books/internal/service"
	"github.com/EmirShimshir/crud-books/internal/transport/rest"
	"github.com/EmirShimshir/crud-books/pkg/database"
	"log"
	"net/http"
	"time"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("config: %+v\n", cfg)

	db, err := database.NewPosgresqlConnection(
		database.ConnectionSettings{
			Host:     cfg.DB.Host,
			Port:     cfg.DB.Port,
			Username: cfg.DB.Username,
			DBName:   cfg.DB.Name,
			SSLMode:  cfg.DB.SSLMode,
			Password: cfg.DB.Password,
		})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	booksRepo := psql.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC822))

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
