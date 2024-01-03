package routes

import "github.com/labstack/echo/v4"

func Load(e *echo.Echo) error {
	albums := e.Group("/albums")

	albums.GET("", getAlbums)

	return nil
}
