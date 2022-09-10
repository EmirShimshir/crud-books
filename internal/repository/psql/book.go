package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/EmirShimshir/crud-books/internal/domain"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db}
}

func (b *BookRepository) Create(ctx context.Context, book domain.Book) (*domain.Book, error) {
	query := "insert into books (title, author, publish_date, rating) values ($1, $2, $3, $4) returning id"
	err := b.db.QueryRowContext(ctx, query, book.Title, book.Author, book.PublishDate, book.Rating).
		Scan(&book.ID)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *BookRepository) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	var book domain.Book
	query := "select id, title, author, publish_date, rating from books where id=$1"
	err := b.db.QueryRowContext(ctx, query, id).
		Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrBookNotFound
	}
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *BookRepository) List(ctx context.Context) ([]domain.Book, error) {
	query := "select id, title, author, publish_date, rating from books"
	rows, err := b.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	books := make([]domain.Book, 0)
	for rows.Next() {
		var book domain.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepository) Delete(ctx context.Context, id int64) error {
	query := "delete from books where id=$1"
	res, err := b.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	countRowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if countRowsAffected == 0 {
		return domain.ErrDeleteFailed
	}

	return nil
}

func (b *BookRepository) Update(ctx context.Context, id int64, inp domain.UpdateBookInput) (*domain.Book, error) {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	if inp.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *inp.Title)
		argId++
	}

	if inp.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, *inp.Author)
		argId++
	}

	if inp.PublishDate != nil {
		setValues = append(setValues, fmt.Sprintf("publish_date=$%d", argId))
		args = append(args, *inp.PublishDate)
		argId++
	}

	if inp.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *inp.Rating)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("update books set %s where id=$%d returning id, title, author, publish_date, rating", setQuery, argId)

	args = append(args, id)

	var book domain.Book
	err := b.db.QueryRowContext(ctx, query, args...).
		Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating)

	return &book, err
}
