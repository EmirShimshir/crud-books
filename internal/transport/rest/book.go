package rest

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/EmirShimshir/crud-books/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initBook(router *gin.RouterGroup) {
	book := router.Group("/book")
	{
		book.GET("/", h.getAllBooks)
		book.POST("/", h.createBook)
		book.GET("/:id", h.getBookByID)
		book.PUT("/:id", h.updateBookByID)
		book.DELETE("/:id", h.deleteBookByID)
	}
}

func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.services.GetBookService().List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "getAllBooks()", "getting books from db", err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) createBook(c *gin.Context) {
	book := new(domain.Book)

	err := c.ShouldBindJSON(book)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "createBook()", "reading request body", err.Error())
		return
	}

	book, err = h.services.GetBookService().Create(context.TODO(), *book)
	if err != nil {
		newErrorResponse(c,
			http.StatusInternalServerError, "createBook()", "adding book to db", err.Error())
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *Handler) getBookByID(c *gin.Context) {
	id, err := getIdFromRequest(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "getBookByID()", "getting ID from request", err.Error())
		return
	}

	book, err := h.services.GetBookService().GetByID(c.Request.Context(), id)
	if err != nil {
		var statusCode int

		switch {
		case errors.Is(err, domain.ErrBookNotFound):
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}

		newErrorResponse(c, statusCode, "getBookByID()", "getting book from db", err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) updateBookByID(c *gin.Context) {
	id, err := getIdFromRequest(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "updateBookByID()", "getting ID from request", err.Error())
		return
	}

	updateBook := new(domain.UpdateBookInput)

	err = c.ShouldBindJSON(updateBook)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "updateBookByID()", "reading request body", err.Error())
		return
	}

	book, err := h.services.GetBookService().Update(c.Request.Context(), id, *updateBook)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "updateBookByID()", "updating book in db", err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) deleteBookByID(c *gin.Context) {
	id, err := getIdFromRequest(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "deleteBookByID()", "getting ID from request", err.Error())
		return
	}

	err = h.services.GetBookService().Delete(c.Request.Context(), id)
	if err != nil {
		var statusCode int

		switch {
		case errors.Is(err, domain.ErrDeleteFailed):
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}

		newErrorResponse(c, statusCode, "deleteBookByID()", "deleting book from db", err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"OK"})
}

func getIdFromRequest(c *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}

	if id < 1 {
		return 0, domain.ErrInvalidID
	}

	return id, nil
}
