package main

import (
	"os"

	"yuzelabs/src/config"
	"yuzelabs/src/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.LoadEnv()
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	db := config.ConnectDatabase()

	defer db.Close()

	router.Load(e, db)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
