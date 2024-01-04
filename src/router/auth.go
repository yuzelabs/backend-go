package router

import (
	"yuzelabs/src/controller"

	"github.com/labstack/echo/v4"
)

func auth(e *echo.Echo) error {
	router := e.Group("/auth")

	router.GET("/nonce", controller.AuthGenerateNonceController)

	return nil
}
