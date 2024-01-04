package controller

import (
	"net/http"
	"yuzelabs/src/service"

	"github.com/labstack/echo/v4"
)

func AuthGenerateNonceController(c echo.Context) error {
	nonce := service.AuthGenerateNonceUseCase()

	data := map[string]string{"nonce": nonce}

	return c.JSON(http.StatusOK, data)
}
