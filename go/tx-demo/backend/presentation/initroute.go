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

	e.POST("/item", itemHandler.Insert())
	e.GET("/item/:id", itemHandler.Get())
	e.POST("/item/title", itemHandler.UpdateTitle())
	e.POST("/item/content", itemHandler.UpdateContent())

	e.POST("/sample", sampleHandler.UpdateNameAndTitle())

	e.Logger.Fatal(e.Start(":8888"))
}
