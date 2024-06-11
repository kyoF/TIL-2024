package presentation

import (
	"backend/application"

	"github.com/labstack/echo/v4"
)

func InitRoute(router IRouter) {
	e := echo.New()

	e.GET("/users", router.GetUsers())

	e.Logger.Fatal(e.Start(":5000"))
}

type IRouter interface {
	GetUsers() echo.HandlerFunc
}

type router struct {
	usecase application.IUsecase
}

func NewRouter(usecase application.IUsecase) IRouter {
	return &router{
		usecase: usecase,
	}
}

func (r *router) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
