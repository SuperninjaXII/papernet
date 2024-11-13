package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Link        string `json:"link"`
	Cartegory1  string `json:"cartegory1"`
	Cartegory2  string `json:"cartegory2"`
}
