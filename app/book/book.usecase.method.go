package book

import (
	"context"
	"go-echo-gorm-tempate/domain"
)

func (b *usecase) AddBooks(ctx context.Context, books []domain.Book) error {
	return b.bookRepo.AddBooks(ctx, books)
}

func (b *usecase) FindAll(ctx context.Context) ([]domain.Book, error) {
	return b.bookRepo.FindAll(ctx)
}
