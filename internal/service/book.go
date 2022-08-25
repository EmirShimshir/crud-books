package service

import (
	"context"
	"time"

	"github.com/EmirShimshir/crud-books/internal/domain"
)

type BookService struct {
	repo domain.BookRepository
}

func NewBookService(repo domain.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (b *BookService) Create(ctx context.Context, book domain.Book) error {
	if book.PublishDate.IsZero() {
		book.PublishDate = time.Now()
	}

	return b.repo.Create(ctx, book)
}

func (b *BookService) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	return b.repo.GetByID(ctx, id)
}

func (b *BookService) List(ctx context.Context) ([]domain.Book, error) {
	return b.repo.List(ctx)
}

func (b *BookService) Delete(ctx context.Context, id int64) error {
	return b.repo.Delete(ctx, id)
}

func (b *BookService) Update(ctx context.Context, id int64, inp domain.UpdateBookInput) error {
	return b.repo.Update(ctx, id, inp)
}
