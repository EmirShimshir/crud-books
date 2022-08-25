package service

import (
	"context"
	"time"

	"github.com/EmirShimshir/crud-books/internal/domain"
)

type Books struct {
	repo domain.BookRepository
}

func NewBooks(repo domain.BookRepository) *Books {
	return &Books{
		repo: repo,
	}
}

func (b *Books) Create(ctx context.Context, book domain.Book) error {
	if book.PublishDate.IsZero() {
		book.PublishDate = time.Now()
	}

	return b.repo.Create(ctx, book)
}

func (b *Books) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	return b.repo.GetByID(ctx, id)
}

func (b *Books) List(ctx context.Context) ([]domain.Book, error) {
	return b.repo.List(ctx)
}

func (b *Books) Delete(ctx context.Context, id int64) error {
	return b.repo.Delete(ctx, id)
}

func (b *Books) Update(ctx context.Context, id int64, inp domain.UpdateBookInput) error {
	return b.repo.Update(ctx, id, inp)
}
