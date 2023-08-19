package author

import (
	"context"
	"go-echo-gorm-tempate/domain"
	"sync"

	"gorm.io/gorm"
)

var (
	authorRepository     domain.AuthorRepository
	authorRepositoryOnce sync.Once
)

type Repository struct {
	model *gorm.DB
}

// NewAuthorRepository ...
func NewAuthorRepository(db *gorm.DB) domain.AuthorRepository {
	authorRepositoryOnce.Do(func() {
		authorRepository = &Repository{
			model: db.Model(&domain.Author{}),
		}
	})

	return authorRepository
}

func (r *Repository) FindAll(ctx context.Context) (authors []domain.Author, err error) {
	authors = []domain.Author{}

	err = r.model.
		WithContext(ctx).
		Find(&authors).
		Error

	return
}

func (r *Repository) FindByID(ctx context.Context, authorID uint64) (*domain.Author, error) {
	author := new(domain.Author)

	err := r.model.
		WithContext(ctx).
		First(author, authorID).
		Error

	if err != nil {
		return nil, err
	}

	return author, nil
}
