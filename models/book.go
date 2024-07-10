package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	Link string `json:"link"`
	Image string `json:"image"`
}