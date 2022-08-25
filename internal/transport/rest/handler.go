package rest

import (
	"github.com/EmirShimshir/crud-books/internal/domain"
	"github.com/gorilla/mux"
	"net/http"
)

type Services interface {
	GetBookService() domain.BookService
}
type Handler struct {
	services Services
}

func NewHandler(services Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRouter() http.Handler {
	router := mux.NewRouter()

	router.Use(loggerMiddleware)

	h.initBook(router)

	return router
}
