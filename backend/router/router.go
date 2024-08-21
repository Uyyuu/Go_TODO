package router

import (
	"go-rest-api/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.GET("/", connect)
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	return e
}

func connect(c echo.Context) error {
	return c.String(http.StatusOK, "Success!")
}