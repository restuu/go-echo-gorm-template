package repository

import (
	"context"
	"go-echo-gorm-tempate/pkg/author/model"
	"sync"

	"gorm.io/gorm"
)

// AuthorRepository ...
type AuthorRepository interface {
	FindAll(context.Context) ([]model.Author, error)
	FindByID(ctx context.Context, authorID uint64) (*model.Author, error)
}

var (
	_authorRepository   AuthorRepository
	_authorRepositoryOnce sync.Once
)

type authorRepository struct {
	model *gorm.DB
}

func (r *authorRepository) FindAll(ctx context.Context) (authors []model.Author, err error) {
	authors = []model.Author{}

	err = r.model.
		WithContext(ctx).
		Find(&authors).
		Error

	return
}

func (r *authorRepository) FindByID(ctx context.Context, authorID uint64) (*model.Author, error) {
	author := new(model.Author)

	err := r.model.
		WithContext(ctx).
		First(author, authorID).
		Error

	if err != nil {
		return nil, err
	}

	return author, nil
}

// NewUserRepository ...
func NewUserRepository(db *gorm.DB) AuthorRepository {
	_authorRepositoryOnce.Do(func() {
		_authorRepository = &authorRepository{
			model: db.Model(&model.Author{}),
		}
	})

	return _authorRepository
}
