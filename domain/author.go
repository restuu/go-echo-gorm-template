package domain

import (
	"context"
	"time"
)

// Author ...
type Author struct {
	ID        uint64    `json:"id" gorm:"primaryKey" `
	FullName  string    `json:"full_name" gorm:"not null" `
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" `
	CreatedBy string    `json:"created_by" gorm:"not null" `
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" `
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index" `
	DeletedBy string    `json:"deleted_by"`
}

type AuthorUsecase interface {
	FindAll(ctx context.Context) ([]Author, error)
}

// AuthorRepository ...
type AuthorRepository interface {
	FindAll(context.Context) ([]Author, error)
	FindByID(ctx context.Context, authorID uint64) (*Author, error)
}
