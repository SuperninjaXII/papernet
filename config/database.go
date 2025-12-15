package config

import (
	"fmt"
	"os"
	"papernet/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s  dbname=%s port=%s", os.Getenv("DB_ADDRESS"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// Try to connect to Postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to Postgres database")
	}

	// Auto-migrate the models
	DB.AutoMigrate(&model.Book{}, &model.User{})
}
