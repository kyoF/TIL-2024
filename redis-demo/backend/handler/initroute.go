package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(userRouter IUserHandler) {
    e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    e.POST("/login", userRouter.Login())
    e.POST("/logout", userRouter.Logout())
    e.POST("/signup", userRouter.Signup())
    e.GET("/auth", userRouter.Test())

	e.Logger.Fatal(e.Start(":8888"))
}
