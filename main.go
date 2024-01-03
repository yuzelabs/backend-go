package main

import (
	"os"
	"yuzelabs/routes"
	"yuzelabs/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	configs.load() 
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	routes.load(e)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
