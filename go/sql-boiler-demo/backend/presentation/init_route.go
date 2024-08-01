package presentation

import "github.com/labstack/echo/v4"

func InitRoute(router IRouter) {
	e := echo.New()

	e.GET("/users", router.GetUserProfiles())

	e.Logger.Fatal(e.Start(":5000"))
}
