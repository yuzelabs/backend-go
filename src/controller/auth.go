package controller

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"time"
	"yuzelabs/src/service"

	"github.com/labstack/echo/v4"
	"github.com/spruceid/siwe-go"
)

var cookieName = "nonce"

func AuthGenerateNonceController(c echo.Context) error {
	nonce := service.AuthGenerateNonceUseCase()

	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Path = "/"
	cookie.Value = nonce
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	data := map[string]string{"nonce": nonce}

	return c.JSON(http.StatusOK, data)
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
		data := map[string]string{"message": "Nonce not found"}

		return c.JSON(http.StatusNotFound, data)
	}

	var message *siwe.Message
	var errSiwe error
	message, errSiwe = siwe.ParseMessage(payload.Message)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errSiwe.Error())
	}

	_, errValid := message.ValidNow()

	if errValid != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errValid.Error())
	}

	nonce := &cookie.Value

	var publicKey *ecdsa.PublicKey
	publicKey, errSiwe = message.Verify(payload.Signature, &payload.Domain, nonce, nil)

	if errSiwe != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errSiwe.Error())
	}

	fmt.Println(publicKey)

	return c.JSON(http.StatusOK, map[string]string{"data": "Ok"})
}
