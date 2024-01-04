package controller

import (
	"net/http"
	"time"
	"yuzelabs/src/service"

	"github.com/labstack/echo/v4"
)

var cookieName = "@yuze-nonce"

func AuthGenerateNonceController(c echo.Context) error {
	nonce := service.AuthGenerateNonceUseCase()

	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = nonce
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	data := map[string]string{"nonce": nonce}

	return c.JSON(http.StatusOK, data)
}

func AuthVerifyMessageController(c echo.Context) error {
	cookie, err := c.Cookie(cookieName)

	if err != nil {
		data := map[string]string{"message": "Nonce not found"}

		return c.JSON(http.StatusNotFound, data)
	}

	data := map[string]string{"nonce": cookie.Value}

	return c.JSON(http.StatusOK, data)
}
