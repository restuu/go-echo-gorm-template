package model

import "time"

// Author ...
type Author struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	FullName  string    `gorm:"not null" json:"full_name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy string    `gorm:"not null" json:"created_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	DeletedBy string    `json:"deleted_by"`
}
