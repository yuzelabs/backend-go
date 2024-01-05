package main

import (
	"net/http"
	"os"

	"yuzelabs/src/config"
	"yuzelabs/src/router"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func init() {
	config.LoadEnv()
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))

	db := config.ConnectDatabase()

	defer db.Close()

	e.Validator = &CustomValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	router.Load(e, db)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
