package config

import (
  "gorm.io/gorm"
  "github.com/glebarez/sqlite"
  "search/models"
)

var DB *gorm.DB

func Database() {
  var err error
  DB, err = gorm.Open(sqlite.Open("admin.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  DB.AutoMigrate(&models.Book{})
}