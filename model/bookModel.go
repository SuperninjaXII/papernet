package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Image       string         `form:"image"`
	File        string         `form:"file"`
	Cartegories string         `json:"Cartegories"`
	Author      string         `json:"author"`
	CreatedAt   time.Time      `json:"created_at"`              // Creation timestamp
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Soft delete field
}
