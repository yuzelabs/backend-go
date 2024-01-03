package main

import (
	"os"

	"yuzelabs/src/config"
	"yuzelabs/src/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.Init()
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	routes.Load(e)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
