package handler

import (
	"backend/application/dto"
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User interface {
	Get() echo.HandlerFunc
	Insert() echo.HandlerFunc
	UpdateName() echo.HandlerFunc
	UpdateAge() echo.HandlerFunc
	Test() echo.HandlerFunc
}

type user struct {
	usecase usecase.User
}

func NewUser(usecase usecase.User) User {
	return &user{usecase: usecase}
}

func (u *user) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := "kyofsampke"
		user, err := u.usecase.Get(userId)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not get user"},
			)
		}
		return c.JSON(http.StatusOK, user)
	}
}
func (u *user) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.User)
		if err := c.Bind(req); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				map[string]string{"message": "Invalid request body"},
			)
		}
		err := u.usecase.Insert(req.UserId, req.Name, req.Age)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not create user"},
			)
		}
		return c.JSON(
			http.StatusOK,
			map[string]string{"message": "Success create user"},
		)
	}
}

func (u *user) UpdateName() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.User)
		if err := c.Bind(req); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				map[string]string{"message": "Invalid request body"},
			)
		}
		err := u.usecase.UpdateName(req.UserId, req.Name)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not update user"},
			)
		}
		return c.JSON(
			http.StatusOK,
			map[string]string{"message": "Success update user"},
		)
	}
}

func (u *user) UpdateAge() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.User)
		if err := c.Bind(req); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				map[string]string{"message": "Invalid request body"},
			)
		}
		err := u.usecase.UpdateAge(req.UserId, req.Age)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not update user"},
			)
		}
		return c.JSON(
			http.StatusOK,
			map[string]string{"message": "Success update user"},
		)
	}
}

func (u *user) Test() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello World!"})
	}
}
