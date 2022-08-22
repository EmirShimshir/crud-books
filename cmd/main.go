package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/EmirShimshir/crud-books/internal/config"
	"github.com/EmirShimshir/crud-books/internal/repository/psql"
	"github.com/EmirShimshir/crud-books/internal/service"
	"github.com/EmirShimshir/crud-books/internal/transport/rest"
	"github.com/EmirShimshir/crud-books/pkg/database"

	log "github.com/sirupsen/logrus"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

// todo
// swagger
// gin?

func main() {
	// init config
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	// init db
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

	// init deps
	booksRepo := psql.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	// init server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	log.Info("SERVER STARTED")

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
