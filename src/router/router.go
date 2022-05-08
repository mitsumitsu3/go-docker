package router

import (
	"net/http"
	"www/src/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})
	e.GET("/user/:id", controller.GetUser)
	e.POST("/user", controller.CreateUser)
	e.PUT("/user/:id", controller.UpdateUser)
	e.DELETE("/user/:id", controller.DeleteUser)

	return e
}
