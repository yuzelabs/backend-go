package routes

import (
	_ "yuzelabs/src/docs"

	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
)

func Load(e *echo.Echo) error {
	e.GET("/docs/*", swagger.WrapHandler)

	albums := e.Group("/albums")

	albums.GET("", getAlbums)

	return nil
}
