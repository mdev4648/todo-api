package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Title       string `gorm:"size:200;not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Completed   bool   `gorm:"default:false" json:"completed"`

	UserID uint `gorm:"not null;index" json:"user_id"`
}
