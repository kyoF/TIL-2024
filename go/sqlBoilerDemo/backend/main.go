package main

import (
	"backend/presentation"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/user", presentation.GetUser())
}
