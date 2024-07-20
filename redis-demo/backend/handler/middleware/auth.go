package middleware

import (
	"backend/application/queryservice"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IMiddleware interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}

type middleware struct {
	queryService queryservice.IAuthQueriService
}

func NewMiddleware(queryService queryservice.IAuthQueriService) IMiddleware {
	return &middleware{queryService: queryService}
}

func (m *middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_id")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing session ID"})
		}

		sessionId := cookie.Value
		name, err := m.queryService.Get(sessionId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking session ID"})
		}

		c.Set("username", name)
		return next(c)
	}
}
