package test

import (
	"log"
	"os"
	"testing"

	"github.com/StefanMoller1/go_library/models"
)

func TestMain(m *testing.M) {
	log := log.Default()
	setupTestServer(log)
	exitCode := m.Run()
	TearDownTestServer(log)
	os.Exit(exitCode)
}

// func TestLibraryAPI(t *testing.T) {
// 	setupTestServer()
// 	id := CreateBookTest(t)
// 	GetBookTest(t, id)

// }

func TestCreateBook(t *testing.T) {
	id, err := CreateBook(models.Book{Name: "Test Book", Description: "Some book to Test With", Author: "Test Author"})
	if err != nil {
		t.Fatalf("client: error creating book: %s\n", err)
	}
	t.Logf("ID Created: %s", id)
}

func TestGetBook(t *testing.T) {
	book := models.Book{Name: "Test Book", Description: "Some book to Test With", Author: "Test Author"}
	id, err := CreateBook(book)
	if err != nil {
		t.Fatalf("client: error creating book: %s\n", err)
	}

	data, err := GetBook(id)
	if err != nil {
		t.Fatalf("client: error getting book: %s\n", err)
	}

	if data.ID != id {
		t.Fatalf("client: expected id %s, got %s", id, data.ID)
	}

	if data.Name != book.Name {
		t.Fatalf("client: expected name %s, got %s", book.Name, data.Name)
	}

	if data.Description != book.Description {
		t.Fatalf("client: expected description %s, got %s", book.Description, data.Description)
	}

	if data.Author != book.Author {
		t.Fatalf("client: expected author %s, got %s", book.Author, data.Author)
	}
}

func TestGetBooks(t *testing.T) {
	book := models.Book{Name: "Test Book", Description: "Some book to Test With", Author: "Test Author"}
	_, err := CreateBook(book)
	if err != nil {
		t.Fatalf("client: error creating book: %s\n", err)
	}

	book = models.Book{Name: "Test Book2", Description: "Some book to Test With2", Author: "Test Author2"}
	_, err = CreateBook(book)
	if err != nil {
		t.Fatalf("client: error creating book: %s\n", err)
	}

	data, err := GetBooks()
	if err != nil {
		t.Fatalf("client: error getting book: %s\n", err)
	}

	if len(data) <= 2 {
		t.Fatalf("client: expected at least 2 books, got %d", len(data))
	}
}

func TestUpdateBook(t *testing.T) {
	book := models.Book{Name: "Test Book", Description: "Some book to Test With", Author: "Test Author"}
	id, err := CreateBook(book)
	if err != nil {
		t.Fatalf("client: error creating book: %s\n", err)
	}

	book2 := models.Book{ID: id, Name: "Test Book2", Description: "Some book to Test With2", Author: "Test Author2"}
	err = UpdateBook(book2)
	if err != nil {
		t.Fatalf("client: error updating book: %s\n", err)
	}

	data, err := GetBook(id)
	if err != nil {
		t.Fatalf("client: error getting book: %s\n", err)
	}

	if data.ID != id {
		t.Fatalf("client: expected id %s, got %s", id, data.ID)
	}

	if data.Name != book2.Name {
		t.Fatalf("client: expected name %s, got %s", book2.Name, data.Name)
	}

	if data.Description != book2.Description {
		t.Fatalf("client: expected description %s, got %s", book2.Description, data.Description)
	}

	if data.Author != book2.Author {
		t.Fatalf("client: expected author %s, got %s", book2.Author, data.Author)
	}
}

func TestDeleteBook(t *testing.T) {
	book := models.Book{Name: "Test Book", Description: "Some book to Test With", Author: "Test Author"}
	id, err := CreateBook(book)
	if err != nil {
		t.Fatalf("client: error creating book: %s\n", err)
	}

	err = DeleteBook(id)
	if err != nil {
		t.Fatalf("client: error deleting book: %s\n", err)
	}

	data, err := GetBook(id)
	if err != nil {
		t.Fatalf("client: error getting book: %s\n", err)
	}

	if data.ID != "" {
		t.Fatalf("client: expected nil, got %s", data)
	}
}
