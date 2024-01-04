package routes

import (
	_ "yuzelabs/src/docs"

	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
)

func Load(e *echo.Echo, db *dbr.Connection) error {
	e.GET("/docs/*", swagger.WrapHandler)

	auth(e)

	return nil
}
