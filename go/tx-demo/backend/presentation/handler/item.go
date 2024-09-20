package handler

import (
	"backend/application/dto"
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Item interface {
	Get() echo.HandlerFunc
	Insert() echo.HandlerFunc
	UpdateTitle() echo.HandlerFunc
	UpdateContent() echo.HandlerFunc
}

type item struct {
	usecase usecase.Item
}

func NewItem(usecase usecase.Item) Item {
	return &item{usecase: usecase}
}

func (i *item) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		itemId := "1"
		item, err := i.usecase.Get(itemId)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not get item"},
			)
		}
		return c.JSON(http.StatusOK, item)
	}
}
func (i *item) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.Item)
		if err := c.Bind(req); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				map[string]string{"message": "Invalid request body"},
			)
		}
		err := i.usecase.Insert(req.ItemId, req.Title, req.Content)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not create item"},
			)
		}
		return c.JSON(
			http.StatusOK,
			map[string]string{"message": "Success create item"},
		)
	}
}

func (i *item) UpdateTitle() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.Item)
		if err := c.Bind(req); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				map[string]string{"message": "Invalid request body"},
			)
		}
		err := i.usecase.UpdateTitle(req.ItemId, req.Title)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not update item"},
			)
		}
		return c.JSON(
			http.StatusOK,
			map[string]string{"message": "Success update item"},
		)
	}
}

func (i *item) UpdateContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.Item)
		if err := c.Bind(req); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				map[string]string{"message": "Invalid request body"},
			)
		}
		err := i.usecase.UpdateContent(req.ItemId, req.Content)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not update item"},
			)
		}
		return c.JSON(
			http.StatusOK,
			map[string]string{"message": "Success update item"},
		)
	}
}
