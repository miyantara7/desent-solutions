package store

import "github.com/miyantara7/desent-solutions/internal/shared/model"

type BookRepository interface {
	Create(model.Book) model.Book
	GetAll(author string) []model.Book
	GetByID(id string) (model.Book, bool)
	Update(id string, req model.Book) (model.Book, bool)
	Delete(id string) bool
}
