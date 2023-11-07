package app

import "github.com/StefanMoller1/go_library/models"

type Library interface {
	SelectAll(pagination *models.Pagination) ([]*models.Book, error)
	Select(id string) (*models.Book, error)
	Create(book *models.Book) error
	Update(book *models.Book) error
	Delete(id string) error
}
