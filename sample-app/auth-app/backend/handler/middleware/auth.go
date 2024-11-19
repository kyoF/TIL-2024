package middleware

import (
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IMiddleware interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}

type middleware struct {
	usecase usecase.IUserUsecase
}

func NewMiddleware(usecase usecase.IUserUsecase) IMiddleware {
	return &middleware{usecase: usecase}
}

func (m *middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_id")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing session ID"})
		}

		sessionId := cookie.Value
		name, err := m.usecase.Get(sessionId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking session ID"})
		}

		c.Set("username", name)
		return next(c)
	}
}
