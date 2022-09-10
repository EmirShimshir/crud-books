package main

import (
	"github.com/EmirShimshir/crud-books/internal/app"
)

// todo
// auth

// @title       crud-books API
// @version     1.0
// @description API Server for CRUD application

// @host     localhost:8080
// @BasePath /

func main() {
	app.Run("configs", "main")
}
