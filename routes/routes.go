package routes

import "github.com/labstack/echo/v4"

func load(e *echo.Echo) {
	albums := e.Group("/albums")

	albums.GET("/", getAlbums)
}
