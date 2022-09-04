package rest

import (
	"github.com/EmirShimshir/crud-books/internal/domain"
	"github.com/gin-gonic/gin"
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
	router := gin.New()

	router.Use(logger(), gin.Recovery())

	h.initBook(&router.RouterGroup)

	return router
}
