package domain

import (
	"context"
	"errors"
	"time"
)

var ErrBookNotFound = errors.New("book not found")

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishDate time.Time `json:"publish_date"`
	Rating      int       `json:"rating"`
}

type UpdateBookInput struct {
	Title       *string    `json:"title"`
	Author      *string    `json:"author"`
	PublishDate *time.Time `json:"publish_date"`
	Rating      *int       `json:"rating"`
}

type BookRepository interface {
	Create(ctx context.Context, book Book) error
	GetByID(ctx context.Context, id int64) (Book, error)
	List(ctx context.Context) ([]Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp UpdateBookInput) error
}

type BookService interface {
	Create(ctx context.Context, book Book) error
	GetByID(ctx context.Context, id int64) (Book, error)
	List(ctx context.Context) ([]Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp UpdateBookInput) error
}
