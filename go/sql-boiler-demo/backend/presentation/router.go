package presentation

import (
	"backend/application"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IRouter interface {
	GetUserProfiles() echo.HandlerFunc
}

type router struct {
	usecase application.IUsecase
}

func NewRouter(usecase application.IUsecase) IRouter {
	return &router{
		usecase: usecase,
	}
}

func (r *router) GetUserProfiles() echo.HandlerFunc {
	return func(c echo.Context) error {
		var getUserProfilesDto []GetUsersDto
		profiles, err := r.usecase.GetUserProfiles()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
		}
		for _, profile := range profiles {
			getUserProfilesDto = append(getUserProfilesDto, GetUsersDto{
				UserId:  profile.UserId,
				Name:    profile.Name,
				Profile: profile.Profile,
			})
		}
		return c.JSON(http.StatusOK, getUserProfilesDto)
	}
}
