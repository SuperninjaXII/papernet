package config

import (
	"papernet/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	var err error
	DB, err = gorm.Open(sqlite.Open("admin.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&model.Book{})
}
