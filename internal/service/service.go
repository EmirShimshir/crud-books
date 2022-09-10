package service

import (
	"github.com/EmirShimshir/crud-books/internal/domain"
	"github.com/EmirShimshir/inMemoryCache"
	"time"
)

type Repositories interface {
	GetBookRepository() domain.BookRepository
}

type Services struct {
	bookService *BookService
}

func (ss *Services) GetBookService() domain.BookService {
	return ss.bookService
}

func NewServices(repo Repositories, cache inMemoryCache.Cache, ttl time.Duration) *Services {
	return &Services{
		bookService: NewBookService(repo.GetBookRepository(), cache, ttl),
	}
}
