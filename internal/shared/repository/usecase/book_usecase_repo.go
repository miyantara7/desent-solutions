package usecase

import "github.com/miyantara7/desent-solutions/internal/shared/model"

type BookUsecase interface {
	CreateBook(req model.Book) (model.Book, error)
	GetBooks(author string) []model.Book
	GetBookByID(id string) (model.Book, error)
	UpdateBook(id string, req model.Book) (model.Book, error)
	DeleteBook(id string) error
}
