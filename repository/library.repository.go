package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/StefanMoller1/go_library/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type LibraryRepository struct {
	*pgx.Conn
	log *log.Logger
}

const booksTableName = "books"

func NewLibraryRepository(conn *pgx.Conn, log *log.Logger) *LibraryRepository {
	return &LibraryRepository{conn, log}
}

func (l *LibraryRepository) SelectAll(pagination *models.Pagination) ([]*models.Book, error) {
	var (
		rows pgx.Rows
		s    string
		err  error
	)

	s = strings.Join([]string{"SELECT COUNT(id) FROM", booksTableName}, " ")
	l.Conn.QueryRow(context.Background(), s).Scan(&pagination.Total)

	rows, err = l.Conn.Query(
		context.Background(),
		`SELECT id, name, title, author, description FROM books LIMIT $1 OFFSET $2`,
		pagination.Size,
		pagination.Page,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no books found")
		}

		return nil, err
	}

	defer rows.Close()
	fmt.Println(rows.RawValues())
	var books []*models.Book

	for rows.Next() {
		fmt.Println("here")
		var book models.Book
		err = rows.Scan(
			&book.ID,
			&book.Name,
			&book.Title,
			&book.Author,
			&book.Description,
		)
		fmt.Println("here")
		if err != nil {
			return nil, err
		}

		l.log.Printf("[INFO] selected book %#v", book)

		books = append(books, &book)
	}

	return books, nil
}

func (l *LibraryRepository) Select(id string) (*models.Book, error) {
	var (
		book models.Book
		err  error
	)

	err = l.Conn.QueryRow(
		context.Background(),
		`SELECT id, name, title, author, description FROM books WHERE id = $1`, id).Scan(
		&book.ID,
		&book.Name,
		&book.Title,
		&book.Author,
		&book.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &book, nil
}

func (l *LibraryRepository) Create(book *models.Book) error {
	var (
		err   error
		res   pgconn.CommandTag
		count int64
	)

	s := fmt.Sprintf(`
		INSERT INTO %s (
			id,
			name,
			title,
			author,
			description
		) 
		VALUES ($1, $2, $3, $4, $5)`,
		booksTableName,
	)

	res, err = l.Conn.Exec(context.Background(), s,
		book.ID,
		book.Name,
		book.Title,
		book.Author,
		book.Description,
	)
	if err != nil {
		return err
	}

	count = res.RowsAffected()
	if count != 1 {
		return errors.New("failed to create new book entry")
	}

	return nil
}

func (l *LibraryRepository) Update(book *models.Book) error {
	var (
		err   error
		res   pgconn.CommandTag
		count int64
	)

	s := fmt.Sprintf(`
		UPDATE %s SET 
		name = $2, 
		title = $3, 
		author = $4, 
		description = $5
		WHERE id = $1`,
		booksTableName,
	)

	res, err = l.Conn.Exec(context.Background(), s,
		book.ID,
		book.Name,
		book.Title,
		book.Author,
		book.Description,
	)
	if err != nil {
		return err
	}

	count = res.RowsAffected()
	if count != 1 {
		return errors.New("failed to update book entry")
	}

	return nil
}

func (l *LibraryRepository) Delete(id string) error {
	var (
		err   error
		res   pgconn.CommandTag
		count int64
	)

	s := strings.Join([]string{"DELETE FROM", booksTableName, "WHERE id = $1"}, " ")
	fmt.Println(s)

	res, err = l.Conn.Exec(context.Background(), s, id)
	if err != nil {
		return err
	}

	count = res.RowsAffected()
	if count != 1 {
		return errors.New("failed to delete book entry")
	}

	return nil
}
