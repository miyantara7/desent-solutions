package book

import (
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/miyantara7/desent-solutions/internal/shared/model"
	repo "github.com/miyantara7/desent-solutions/internal/shared/repository/store"
)

type bookRepository struct {
	mu    sync.RWMutex
	books map[string]model.Book
}

func NewBookRepository() repo.BookRepository {
	return &bookRepository{
		books: make(map[string]model.Book),
	}
}

func (r *bookRepository) Create(b model.Book) model.Book {
	r.mu.Lock()
	defer r.mu.Unlock()

	b.ID = uuid.NewString()
	r.books[b.ID] = b
	return b
}

func (r *bookRepository) GetAll(author string) []model.Book {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []model.Book
	for _, b := range r.books {
		if author == "" || strings.Contains(strings.ToLower(b.Author), author) {
			result = append(result, b)
		}
	}
	return result
}

func (r *bookRepository) GetByID(id string) (model.Book, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	b, ok := r.books[id]
	return b, ok
}

func (r *bookRepository) Update(id string, req model.Book) (model.Book, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	b, ok := r.books[id]
	if !ok {
		return model.Book{}, false
	}

	if req.Title != "" {
		b.Title = req.Title
	}
	if req.Author != "" {
		b.Author = req.Author
	}

	r.books[id] = b
	return b, true
}

func (r *bookRepository) Delete(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.books[id]; !ok {
		return false
	}

	delete(r.books, id)
	return true
}
