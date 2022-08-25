package domain

import (
	"context"
	"time"
)

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
	Create(ctx context.Context, book Book) (*Book, error)
	List(ctx context.Context) ([]Book, error)
	GetByID(ctx context.Context, id int64) (*Book, error)
	Update(ctx context.Context, id int64, inp UpdateBookInput) (*Book, error)
	Delete(ctx context.Context, id int64) error
}

type BookService interface {
	Create(ctx context.Context, book Book) (*Book, error)
	List(ctx context.Context) ([]Book, error)
	GetByID(ctx context.Context, id int64) (*Book, error)
	Update(ctx context.Context, id int64, inp UpdateBookInput) (*Book, error)
	Delete(ctx context.Context, id int64) error
}
