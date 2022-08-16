package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/EmirShimshir/crud-books/internal/domain"
)

type Books struct {
	db *sql.DB
}

func NewBooks(db *sql.DB) *Books {
	return &Books{db}
}

func (b *Books) Create(ctx context.Context, book domain.Book) error {
	_, err := b.db.ExecContext(ctx, "insert into books (title, author, publish_date, rating) values ($1, $2, $3, $4)",
		book.Title, book.Author, book.PublishDate, book.Rating)

	return err
}

func (b *Books) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	var book domain.Book
	row := b.db.QueryRowContext(ctx, "select id, title, author, publish_date, rating from books where id=$1", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating)
	if err == sql.ErrNoRows {
		return book, domain.ErrBookNotFound
	}
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *Books) GetAll(ctx context.Context) ([]domain.Book, error) {
	rows, err := b.db.QueryContext(ctx, "select id, title, author, publish_date, rating from books")
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

func (b *Books) Delete(ctx context.Context, id int64) error {
	_, err := b.db.ExecContext(ctx, "delete from books where id=$1", id)

	return err
}

func (b *Books) Update(ctx context.Context, id int64, inp domain.UpdateBookInput) error {
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

	query := fmt.Sprintf("update books set %s where id=$%d", setQuery, argId)

	args = append(args, id)

	_, err := b.db.ExecContext(ctx, query, args...)

	return err
}
