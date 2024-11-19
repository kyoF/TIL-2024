package router

import (
	"backend/application/dto"
	"backend/application/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type User interface {
	Login() echo.HandlerFunc
	Logout() echo.HandlerFunc
	Signup() echo.HandlerFunc
	Test() echo.HandlerFunc
}

type user struct {
	usecase usecase.User
}

func NewUser(usecase usecase.User) User {
	return &user{usecase: usecase}
}

func (h *user) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(dto.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
		}

		sessionId, err := h.usecase.Login(user.Username, user.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
		}

		c.SetCookie(
			&http.Cookie{
				Name:     "session_id",
				Value:    sessionId,
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
			},
		)
		return c.JSON(http.StatusOK, map[string]string{"message": "Login successful"})
	}
}

func (h *user) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_id")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "No session ID found"})
		}

		sessionId := cookie.Value

		err = h.usecase.Logout(sessionId)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not delete session"})
		}

		c.SetCookie(&http.Cookie{
			Name:     "session_id",
			Value:    "",
			Expires:  time.Now().Add(-1 * time.Hour),
			HttpOnly: true,
		})
		return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
	}
}

func (h *user) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(dto.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
		}
		err := h.usecase.Signup(user.Username, user.Password)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not create user"},
			)
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
	}
}

func (h *user) Test() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello World!"})
	}
}
