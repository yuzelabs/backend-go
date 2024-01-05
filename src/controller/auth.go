package controller

import (
	"net/http"
	"yuzelabs/src/service"

	"github.com/labstack/echo/v4"
	"github.com/spruceid/siwe-go"
)

func AuthGenerateNonceController(c echo.Context) error {
	nonce := service.AuthGenerateNonceUseCase()

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

	if err != nil {
		data := map[string]string{"message": "Generate a signature first"}

		return c.JSON(http.StatusBadRequest, data)
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

	nonce := message.GetNonce()
	domain := message.GetDomain()
	address := message.GetAddress()

	_, errSiwe = message.Verify(payload.Signature, &domain, &nonce, nil)

	if errSiwe != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errSiwe.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"data": address.String()})
}
