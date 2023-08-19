package book

import (
	"context"
	"go-echo-gorm-tempate/domain"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	_bookRepository     domain.BookRepository
	_bookRepositoryOnce sync.Once
)

type bookRepository struct {
	model *gorm.DB
}

func NewBookRepository(db *gorm.DB) domain.BookRepository {
	_bookRepositoryOnce.Do(func() {
		_bookRepository = &bookRepository{
			model: db.Model(domain.Book{}),
		}
	})

	return _bookRepository
}

// AddBooks implements BookRepository.
func (r *bookRepository) AddBooks(ctx context.Context, books []domain.Book) error {
	return r.model.WithContext(ctx).CreateInBatches(books, 10).Error
}

// FindAll implements BookRepository.
func (r *bookRepository) FindAll(ctx context.Context) ([]domain.Book, error) {
	books := make([]domain.Book, 0)

	err := r.model.
		WithContext(ctx).
		Preload(clause.Associations).
		Find(&books).
		Error

	return books, err
}

// FindByAuthor implements BookRepository.
func (r *bookRepository) FindByAuthor(ctx context.Context, authorID uint64) ([]domain.Book, error) {
	books := make([]domain.Book, 0)

	err := r.model.
		WithContext(ctx).
		Preload(clause.Associations).
		Where(&domain.Book{AuthorID: authorID}).
		Find(&books).
		Error

	return books, err
}
