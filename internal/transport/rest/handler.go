package rest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/EmirShimshir/crud-books/internal/domain"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Books interface {
	Create(ctx context.Context, book domain.Book) error
	GetByID(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.UpdateBookInput) error
}

type Handler struct {
	booksService Books
}

func NewHandler(books Books) *Handler {
	return &Handler{
		booksService: books,
	}
}

func (h *Handler) InitRouter() http.Handler {
	sm := http.NewServeMux()

	sm.HandleFunc("/books/", loggerMiddleware(h.book))
	sm.HandleFunc("/books/id/", loggerMiddleware(h.id))

	return sm
}

func (h *Handler) book(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllBooks(w, r)
	case http.MethodPost:
		h.createBook(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (h *Handler) id(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getBookByID(w, r)
	case http.MethodPut:
		h.updateBookByID(w, r)
	case http.MethodDelete:
		h.deleteBookByID(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.booksService.GetAll(context.TODO())
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(books)
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	book := domain.Book{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("createBook() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(data, &book)
	if err != nil {
		log.Println("createBook() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Create(context.TODO(), book)
	if err != nil {
		log.Println("createBook() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("getBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	book, err := h.booksService.GetByID(context.TODO(), id)
	if err != nil {
		log.Println("getBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	data, err := json.Marshal(book)
	if err != nil {
		log.Println("getBookByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h *Handler) updateBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("updateBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("updateBookByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	updateBook := domain.UpdateBookInput{}

	err = json.Unmarshal(data, &updateBook)
	if err != nil {
		log.Println("updateBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.booksService.Update(context.TODO(), id, updateBook)
	if err != nil {
		log.Println("updateBookByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("deleteBookByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = h.booksService.Delete(context.TODO(), id)
	if err != nil {
		log.Println("deleteBookByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func getIdFromRequest(r *http.Request) (int64, error) {
	id, err := strconv.Atoi(r.URL.String()[len("/books/id/"):])
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return int64(id), nil
}
