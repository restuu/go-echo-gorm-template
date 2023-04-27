package repository

import (
	"context"
	"go-echo-gorm-tempate/pkg/book/model"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BookRepository ...
type BookRepository interface {
	FindAll(ctx context.Context) ([]model.Book, error)
	FindByAuthor(ctx context.Context, authorID uint64) ([]model.Book, error)
}

var (
	_bookRepository     BookRepository
	_bookRepositoryOnce sync.Once
)

type bookRepository struct {
	model *gorm.DB
}

// FindAll implements BookRepository
func (r *bookRepository) FindAll(ctx context.Context) ([]model.Book, error) {
	books := []model.Book{}

	err := r.model.
		WithContext(ctx).
		Preload(clause.Associations).
		Find(&books).
		Error

	return books, err
}

// FindByAuthor implements BookRepository
func (r *bookRepository) FindByAuthor(ctx context.Context, authorID uint64) ([]model.Book, error) {
	books := []model.Book{}

	err := r.model.
		WithContext(ctx).
		Preload(clause.Associations).
		Where(&model.Book{AuthorID: authorID}).
		Find(&books).
		Error

	return books, err
}

func NewBookRepository(db *gorm.DB) BookRepository {
	_bookRepositoryOnce.Do(func() {
		_bookRepository = &bookRepository{
			model: db.Model(model.Book{}),
		}
	})

	return _bookRepository
}
