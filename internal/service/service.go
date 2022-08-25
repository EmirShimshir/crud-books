package service

import "github.com/EmirShimshir/crud-books/internal/domain"

type Repositories interface {
	GetBookRepository() domain.BookRepository
}

type Services struct {
	bookService *BookService
}

func (ss *Services) GetBookService() domain.BookService {
	return ss.bookService
}

func NewServices(repo Repositories) *Services {
	return &Services{
		bookService: NewBookService(repo.GetBookRepository()),
	}
}
