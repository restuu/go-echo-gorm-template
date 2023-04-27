package service

import (
	"context"
	"go-echo-gorm-tempate/pkg/book/model"
	"go-echo-gorm-tempate/pkg/book/repository"
	"sync"
)

// BookGettingService ...
type BookGettingService interface {
	FindAll(ctx context.Context) ([]model.Book, error)
}

type bookGettingService struct {
	bookRepo repository.BookRepository
}

// FindAll implements BookGettingService
func (s *bookGettingService) FindAll(ctx context.Context) ([]model.Book, error) {
	return s.bookRepo.FindAll(ctx)
}

var (
	_bookGettingService     BookGettingService
	_bookGettingServiceOnce sync.Once
)

// NewBookGettingService instantiate BookGettingService
func NewBookGettingService(bookRepository repository.BookRepository) BookGettingService {
	_bookGettingServiceOnce.Do(func() {
		_bookGettingService = &bookGettingService{
			bookRepo: bookRepository,
		}
	})

	return _bookGettingService
}
