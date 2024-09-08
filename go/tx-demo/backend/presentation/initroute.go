package presentation

import (
	"github.com/labstack/echo/v4"

	"backend/presentation/handler"
)

func InitRoute(userHandler handler.User) {
	e := echo.New()

	e.GET("/user", userHandler.Get())
	e.POST("/user", userHandler.Insert())
	e.POST("/user/name", userHandler.UpdateName())
	e.POST("/user/age", userHandler.UpdateAge())
	e.GET("/test", userHandler.Test())

	e.Logger.Fatal(e.Start(":8888"))
}
