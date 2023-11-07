package router

import (
	"errors"
	"net/http"

	"github.com/StefanMoller1/go_library/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type Book struct {
	*models.Book
}

func (b *Book) Bind(r *http.Request) error {
	return nil
}

func (m *Manager) LibraryRouter(r chi.Router) {
	r.With(m.Paginate).Get("/", m.selectAll)
	r.Get("/{id}", m.selectBook)
	r.Post("/", m.createBook)
	r.Put("/{id}", m.updateBook)
	r.Delete("/{id}", m.deleteBook)
}

func (m *Manager) selectAll(w http.ResponseWriter, r *http.Request) {
	var (
		books      []*models.Book
		pagination *models.Pagination
		ok         bool
		safe       string
		err        error
	)

	defer func() {

		if err != nil {
			m.Log.Printf("[ERROR] selecting all books: %v", err)
			render.Status(r, 500)
			render.JSON(w, r, models.Response{Error: safe})
			return
		}
		render.JSON(w, r, models.Response{Data: books, Pagination: pagination})
	}()

	pagination, ok = r.Context().Value(paginationContext).(*models.Pagination)
	if !ok {
		safe = "failed to parse pagination context"
		err = errors.New(safe)
		return
	}

	books, err = m.Library.SelectAll(pagination)
	if err != nil {
		safe = "failed to select all books"
		return
	}
}

func (m *Manager) selectBook(w http.ResponseWriter, r *http.Request) {
	var (
		book *models.Book
		safe string
		err  error
	)

	defer func() {

		if err != nil {
			m.Log.Printf("[ERROR] selecting book: %v", err)
			render.Status(r, 500)
			render.JSON(w, r, models.Response{Error: safe})
			return
		}
		render.JSON(w, r, models.Response{Data: book})
	}()

	id := chi.URLParam(r, "id")

	book, err = m.Library.Select(id)
	if err != nil {
		safe = "failed to select all books"
		return
	}
}

func (m *Manager) createBook(w http.ResponseWriter, r *http.Request) {
	var (
		book Book
		safe string
		err  error
	)

	defer func() {

		if err != nil {
			m.Log.Printf("[ERROR] creating new books: %v", err)
			render.Status(r, 500)
			render.JSON(w, r, models.Response{Error: safe})
			return
		}

		render.JSON(w, r, models.Response{Data: book})
	}()

	err = render.Bind(r, &book)
	if err != nil {
		safe = "failed to parse create book request"
		return
	}

	book.ID = uuid.NewString()

	err = m.Library.Create(book.Book)
	if err != nil {
		safe = "failed to create book"
	}
}

func (m *Manager) updateBook(w http.ResponseWriter, r *http.Request) {
	var (
		book Book
		safe string
		err  error
	)

	defer func() {

		if err != nil {
			m.Log.Printf("[ERROR] updating book: %v", err)
			render.Status(r, 500)
			render.JSON(w, r, models.Response{Error: safe})
			return
		}
		render.JSON(w, r, models.Response{Data: book})
	}()

	id := chi.URLParam(r, "id")

	err = render.Bind(r, &book)
	if err != nil {
		safe = "failed to parse create book request"
		return
	}

	book.ID = id

	err = m.Library.Update(book.Book)
	if err != nil {
		safe = "failed to update book"
	}
}

func (m *Manager) deleteBook(w http.ResponseWriter, r *http.Request) {
	var (
		safe string
		err  error
	)

	defer func() {

		if err != nil {
			m.Log.Printf("[ERROR] deleting book: %v", err)
			render.Status(r, 500)
			render.JSON(w, r, models.Response{Error: safe})
			return
		}
		render.JSON(w, r, models.Response{Error: safe})
	}()

	id := chi.URLParam(r, "id")

	err = m.Library.Delete(id)
	if err != nil {
		safe = "failed to delete book"
	}
}
