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
}

var (
	_userRepository     AuthorRepository
	_userRepositoryOnce sync.Once
)

type userRepository struct {
	model *gorm.DB
}

func (u *userRepository) FindAll(ctx context.Context) (users []model.Author, err error) {
	err = u.model.Find(&users).Error

	return
}

// NewUserRepository ...
func NewUserRepository(db *gorm.DB) AuthorRepository {
	_userRepositoryOnce.Do(func() {
		_userRepository = &userRepository{
			model: db.Model(&model.Author{}),
		}
	})

	return _userRepository
}
