package config

import (
	"log"
	"os"

	"yuzelabs/src/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func autoMigrate() {
	db.AutoMigrate(&model.User{})
}

func connectDatabase() {
	var err error

	url := os.Getenv("DATABASE_URL")

	db, err = gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
}

func GetDatabase() *gorm.DB {
	if db == nil {
		log.Fatal("Database not initialized")
	}

	return db
}
