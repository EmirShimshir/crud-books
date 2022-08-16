package main

import (
	"github.com/EmirShimshir/crud-books/internal/repository/psql"
	"github.com/EmirShimshir/crud-books/internal/service"
	"github.com/EmirShimshir/crud-books/internal/transport/rest"
	"github.com/EmirShimshir/crud-books/pkg/database"
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := database.NewPosgresqlConnection(
		database.ConnectionSettings{
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			DBName:   "postgres",
			SSLMode:  "disable",
			Password: "qwerty123",
		})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	booksRepo := psql.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	srv := http.Server{
		Addr:    ":10000",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC822))

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
