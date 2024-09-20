package presentation

import (
	"github.com/labstack/echo/v4"

	"backend/presentation/handler"
)

func InitRoute(
	userHandler handler.User,
	itemHandler handler.Item,
	sampleHandler handler.Sample,
) {
	e := echo.New()

	e.GET("/user", userHandler.Get())
	e.POST("/user", userHandler.Insert())
	e.POST("/user/name", userHandler.UpdateName())
	e.POST("/user/age", userHandler.UpdateAge())
	e.GET("/test", userHandler.Test())

	e.GET("/item", userHandler.Get())
	e.POST("/item", userHandler.Insert())
	e.POST("/item/title", userHandler.UpdateName())
	e.POST("/item/content", userHandler.UpdateAge())

	e.POST("/sample", sampleHandler.UpdateNameAndTitle())

	e.Logger.Fatal(e.Start(":8888"))
}
