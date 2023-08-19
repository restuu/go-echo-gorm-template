package domain

import (
	"context"
	"time"
)

// Book ...
type Book struct {
	ID        uint64    `json:"id" gorm:"primaryKey" `
	Title     string    `json:"full_name" gorm:"not null" `
	AuthorID  uint64    `json:"author_id" gorm:"not null" `
	Author    *Author   `json:"author"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" `
	CreatedBy string    `json:"created_by" gorm:"not null" `
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" `
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index" `
	DeletedBy string    `json:"deleted_by"`
}

type BookUsecase interface {
	AddBooks(ctx context.Context, books []Book) error
	FindAll(ctx context.Context) ([]Book, error)
}

// BookRepository ...
type BookRepository interface {
	AddBooks(ctx context.Context, books []Book) error
	FindAll(ctx context.Context) ([]Book, error)
	FindByAuthor(ctx context.Context, authorID uint64) ([]Book, error)
}
