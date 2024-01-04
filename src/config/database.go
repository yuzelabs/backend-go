package config

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
)

func ConnectDatabase() *dbr.Connection {
	var err error

	url := os.Getenv("DATABASE_URL")

	db, err := dbr.Open("mysql", url, nil)

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	return db
}
