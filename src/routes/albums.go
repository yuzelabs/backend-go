package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getAlbums(c echo.Context) error {
	data := map[string]string{"data": "Hello world"}

	return c.JSON(http.StatusOK, data)
}
