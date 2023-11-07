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
	r.Get("/", m.selectAll)
	r.Get("/{id}", m.selectBook)
	r.Post("/", m.createBook)
	r.Put("/{id}", m.updateBook)
	r.Delete("/{id}", m.deleteBook)
}

func (m *Manager) selectAll(w http.ResponseWriter, r *http.Request) {
	var (
		books []*models.Book
		safe  string
		err   error
	)

	defer func() {

		if err != nil {
			m.log.Printf("[ERROR] selecting all books: %v", err)
			render.JSON(w, r, models.Response{Error: safe})
			return
		}
		render.JSON(w, r, models.Response{Data: books})
	}()

	resp := m.db.Find(&books)
	if resp.Error != nil {
		err = resp.Error
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
			m.log.Printf("[ERROR] selecting book: %v", err)
			render.JSON(w, r, models.Response{Error: safe})
			return
		}
		render.JSON(w, r, models.Response{Data: book})
	}()

	id := chi.URLParam(r, "id")

	resp := m.db.First(&book, models.Book{ID: id})
	if resp.Error != nil {
		err = resp.Error
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
			m.log.Printf("[ERROR] creating new books: %v", err)
			render.JSON(w, r, models.Response{Error: safe})
			return
		}
		render.JSON(w, r, models.Response{Data: book})
	}()

	err = render.Bind(r, &book)
	if err != nil {
		m.log.Println(err)
		safe = "failed to parse create book request"
		return
	}

	book.ID = uuid.NewString()

	resp := m.db.Create(&book.Book)
	if resp.Error != nil {
		err = resp.Error
		safe = "failed to create book"
	}

	if resp.RowsAffected == 0 {
		safe = "failed to create book"
		err = errors.New(safe)
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
			m.log.Printf("[ERROR] updating book: %v", err)
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

	resp := m.db.Save(&book.Book)
	if resp.Error != nil {
		err = resp.Error
		safe = "failed to update book"
	}

	if resp.RowsAffected == 0 {
		safe = "failed to update book"
		err = errors.New(safe)
	}
}

func (m *Manager) deleteBook(w http.ResponseWriter, r *http.Request) {
	var (
		safe string
		err  error
	)

	defer func() {

		if err != nil {
			m.log.Printf("[ERROR] deleting all books: %v", err)
			render.JSON(w, r, models.Response{Error: safe})
			return
		}
		render.JSON(w, r, models.Response{Error: safe})
	}()

	id := chi.URLParam(r, "id")

	resp := m.db.Delete(&models.Book{ID: id})
	if resp.Error != nil {
		err = resp.Error
		safe = "failed to delete book"
	}

	if resp.RowsAffected == 0 {
		safe = "failed to delete book"
		err = errors.New(safe)
	}
}
