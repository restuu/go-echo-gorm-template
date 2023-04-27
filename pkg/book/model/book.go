package model

import "time"

// Book ...
type Book struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"full_name"`
	AuthorID  uint64    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy string    `gorm:"not null" json:"created_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	DeletedBy string    `json:"deleted_by"`
}
