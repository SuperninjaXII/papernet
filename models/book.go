package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Cartegory   string `json:"cartegory"`
	Link        string `json:"link"`
	Image       string `json:"image"`
}

