package rest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/EmirShimshir/crud-books/internal/domain"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	r.Use(loggerMiddleware)

	booksRouter := r.PathPrefix("/books").Subrouter()
	{
		booksRouter.HandleFunc("/", h.getAllBooks).Methods(http.MethodGet)
		booksRouter.HandleFunc("/", h.createBook).Methods(http.MethodPost)
		booksRouter.HandleFunc("/{id:[0-9]+}", h.getBookByID).Methods(http.MethodGet)
		booksRouter.HandleFunc("/{id:[0-9]+}", h.updateBookByID).Methods(http.MethodPut)
		booksRouter.HandleFunc("/{id:[0-9]+}", h.deleteBookByID).Methods(http.MethodDelete)
	}

	return r
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("getBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := h.booksService.GetByID(context.TODO(), id)
	if err != nil {
		log.Println("getBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(book)
	if err != nil {
		log.Println("getBookByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h *Handler) updateBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("updateBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("updateBookByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updateBook := domain.UpdateBookInput{}

	err = json.Unmarshal(data, &updateBook)
	if err != nil {
		log.Println("updateBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Update(context.TODO(), id, updateBook)
	if err != nil {
		log.Println("updateBookByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("deleteBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Delete(context.TODO(), id)
	if err != nil {
		println(1)
		log.Println("deleteBookByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return id, nil
}
