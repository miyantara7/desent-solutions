package book

import (
	"errors"

	"github.com/miyantara7/desent-solutions/internal/shared/model"
	repo "github.com/miyantara7/desent-solutions/internal/shared/repository/store"
	repoUc "github.com/miyantara7/desent-solutions/internal/shared/repository/usecase"
)

type bookUsecase struct {
	repo repo.BookRepository
}

func NewBookUsecase(r repo.BookRepository) repoUc.BookUsecase {
	return &bookUsecase{repo: r}
}

func (u *bookUsecase) CreateBook(req model.Book) (model.Book, error) {
	if req.Title == "" {
		return model.Book{}, errors.New("title required")
	}
	return u.repo.Create(req), nil
}

func (u *bookUsecase) GetBooks(author string) []model.Book {
	return u.repo.GetAll(author)
}

func (u *bookUsecase) GetBookByID(id string) (model.Book, error) {
	b, ok := u.repo.GetByID(id)
	if !ok {
		return model.Book{}, errors.New("book not found")
	}
	return b, nil
}

func (u *bookUsecase) UpdateBook(id string, req model.Book) (model.Book, error) {
	b, ok := u.repo.Update(id, req)
	if !ok {
		return model.Book{}, errors.New("book not found")
	}
	return b, nil
}

func (u *bookUsecase) DeleteBook(id string) error {
	if !u.repo.Delete(id) {
		return errors.New("book not found")
	}
	return nil
}
