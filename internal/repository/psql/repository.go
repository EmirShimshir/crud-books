package psql

import (
	"database/sql"
	"github.com/EmirShimshir/crud-books/internal/domain"
)

type Repositories struct {
	bookRepository *BookRepository
}

func (rs *Repositories) GetBookRepository() domain.BookRepository {
	return rs.bookRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		bookRepository: NewBookRepository(db),
	}
}
