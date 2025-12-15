package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserName  string         `json:"username"`
	Email     string         `json:"email"`
	Number    string         `json:"number"`
	Password  string         `json:"password"`
}
