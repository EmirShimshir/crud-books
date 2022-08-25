package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/EmirShimshir/crud-books/internal/domain"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) initBook(router *mux.Router) {
	booksRouter := router.PathPrefix("/books").Subrouter()
	{
		booksRouter.HandleFunc("", h.getAllBooks).Methods(http.MethodGet)
		booksRouter.HandleFunc("", h.createBook).Methods(http.MethodPost)
		booksRouter.HandleFunc("/{id}", h.getBookByID).Methods(http.MethodGet)
		booksRouter.HandleFunc("/{id}", h.updateBookByID).Methods(http.MethodPut)
		booksRouter.HandleFunc("/{id}", h.deleteBookByID).Methods(http.MethodDelete)
	}
}

func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	books, err := h.services.GetBookService().List(context.TODO())
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "getAllBooks",
			"problem": "getting books from db",
		}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "getAllBooks",
			"problem": "marshalling books",
		}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	book := domain.Book{}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "createBook",
			"problem": "reading request body",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.services.GetBookService().Create(context.TODO(), book)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "createBook",
			"problem": "adding book to db",
		}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := getIdFromRequest(r)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "getBookByID",
			"problem": "getting ID from request",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := h.services.GetBookService().GetByID(context.TODO(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "getBookByID",
			"problem": "getting book from db",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "getBookByID",
			"problem": "marshalling book",
		}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) updateBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "updateBookByID",
			"problem": "getting ID from request",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateBook := domain.UpdateBookInput{}

	err = json.NewDecoder(r.Body).Decode(&updateBook)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "updateBookByID",
			"problem": "reading request body",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.services.GetBookService().Update(context.TODO(), id, updateBook)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "updateBookByID",
			"problem": "updating book in db",
		}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) deleteBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "deleteBookByID",
			"problem": "getting ID from request",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.services.GetBookService().Delete(context.TODO(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "deleteBookByID",
			"problem": "deleting book from db",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
