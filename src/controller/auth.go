package controller

import (
	"net/http"
	"yuzelabs/src/service"

	"github.com/labstack/echo/v4"
)

var cookieName = "nonce"

func AuthGenerateNonceController(c echo.Context) error {
	nonce := service.AuthGenerateNonceUseCase()

	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = nonce

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"nonce": nonce})
}

func AuthVerifyMessageController(c echo.Context) (err error) {
	payload := new(AuthVerifyMessageModel)

	if err = c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(payload); err != nil {
		return err
	}

	cookie, err := c.Cookie(cookieName)

	if err != nil {
		data := map[string]string{"message": "Generate a signature first"}

		return c.JSON(http.StatusBadRequest, data)
	}

	err = service.AuthVerifyMessageUseCase(payload.Signature, payload.Message, &cookie.Value)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"data": "Ok"})
}
