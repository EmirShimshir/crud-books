package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"

	"github.com/EmirShimshir/crud-books/internal/domain"

	"github.com/EmirShimshir/inMemoryCache"
)

var bookSalt = "_book"
var listSalt = "_list"
var listID int64 = 0

type BookService struct {
	repo  domain.BookRepository
	cache inMemoryCache.Cache
	ttl   time.Duration
}

func NewBookService(repo domain.BookRepository, cache inMemoryCache.Cache, ttl time.Duration) *BookService {
	return &BookService{
		repo:  repo,
		cache: cache,
		ttl:   ttl,
	}
}

func (b *BookService) Create(ctx context.Context, bookInput domain.Book) (*domain.Book, error) {
	if bookInput.PublishDate.IsZero() {
		bookInput.PublishDate = time.Now()
	}

	book, err := b.repo.Create(ctx, bookInput)
	if err == nil {
		_ = b.cache.Set(idToString(book.ID, bookSalt), book, b.ttl)
		listID++
	}

	return book, err
}

func (b *BookService) GetByID(ctx context.Context, id int64) (book *domain.Book, err error) {
	i, err := b.cache.Get(idToString(id, bookSalt))
	var ok bool
	if err == nil {
		log.WithFields(log.Fields{
			"from": "BookService.GetById()",
		}).Debug("Get book from cache")
		book, ok = i.(*domain.Book)
	}

	if !ok {
		log.WithFields(log.Fields{
			"from": "BookService.GetById()",
		}).Debug("Get book from repo")
		book, err = b.repo.GetByID(ctx, id)
	}

	if err == nil {
		_ = b.cache.Set(idToString(id, bookSalt), book, b.ttl)
	}

	return book, err
}

func (b *BookService) List(ctx context.Context) (books []domain.Book, err error) {
	i, err := b.cache.Get(idToString(listID, listSalt))
	var ok bool
	if err == nil {
		log.WithFields(log.Fields{
			"from": "BookService.GetById()",
		}).Debug("Get book from cache")
		books, ok = i.([]domain.Book)
	}

	if !ok {
		log.WithFields(log.Fields{
			"from": "BookService.GetById()",
		}).Debug("Get book from repo")
		books, err = b.repo.List(ctx)
	}

	if err == nil {
		_ = b.cache.Set(idToString(listID, listSalt), books, b.ttl)
	}

	return books, err
}

func (b *BookService) Delete(ctx context.Context, id int64) error {

	err := b.repo.Delete(ctx, id)
	if err == nil {
		_ = b.cache.Delete(idToString(id, bookSalt))
		listID++
	}

	return err
}

func (b *BookService) Update(ctx context.Context, id int64, inp domain.UpdateBookInput) (*domain.Book, error) {
	book, err := b.repo.Update(ctx, id, inp)
	if err == nil {
		_ = b.cache.Set(idToString(book.ID, bookSalt), book, b.ttl)
		listID++
	}

	return book, err
}

func idToString(id int64, cacheSalt string) string {
	return strconv.FormatInt(id, 10) + cacheSalt
}
