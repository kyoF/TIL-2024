package handler

import (
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Sample interface {
	UpdateNameAndTitle() echo.HandlerFunc
}

type sample struct {
	usecase usecase.Sample
}

func NewSample(usecase usecase.Sample) Sample {
	return &sample{usecase: usecase}
}

func (s *sample) UpdateNameAndTitle() echo.HandlerFunc {
	return func(c echo.Context) error {
		itemId := "item1"
		userId := "user1"
		err := s.usecase.UpdateTitleAndName(userId, itemId, "username", "itemtitle")
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not update user and item"},
			)
		}
		return c.JSON(
			http.StatusOK,
			map[string]string{"message": "Success update user and item"},
		)
	}
}
