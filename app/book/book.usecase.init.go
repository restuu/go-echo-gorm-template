package book

import (
	"go-echo-gorm-tempate/domain"
	"sync"
)

type usecase struct {
	bookRepo domain.BookRepository
}

var (
	uc     domain.BookUsecase
	ucOnce sync.Once
)

func NewBookUsecase(bookRepo domain.BookRepository) domain.BookUsecase {
	ucOnce.Do(func() {
		uc = &usecase{bookRepo: bookRepo}
	})

	return uc
}
