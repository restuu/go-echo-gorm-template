package service

import (
	"context"
	"go-echo-gorm-tempate/pkg/author/model"
	"go-echo-gorm-tempate/pkg/author/repository"
	"sync"
)

type AuthorGettingService interface {
	FindAll(ctx context.Context) ([]model.Author, error)
}

var (
	_authorGettingService     AuthorGettingService
	_authorGettingServiceOnce sync.Once
)

type authorGettingService struct {
	authorRepository repository.AuthorRepository
}

// FindAll implements AuthorGettingService.
func (svc *authorGettingService) FindAll(ctx context.Context) ([]model.Author, error) {
	return svc.authorRepository.FindAll(ctx)
}

func NewUserGettingService(authorRepository repository.AuthorRepository) AuthorGettingService {
	_authorGettingServiceOnce.Do(func() {
		_authorGettingService = &authorGettingService{
			authorRepository: authorRepository,
		}
	})

	return _authorGettingService
}
