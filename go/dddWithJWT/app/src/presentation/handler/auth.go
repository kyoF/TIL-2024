package handler

import (
	"app/src/application/dto"
	"app/src/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	Signup() echo.HandlerFunc
	Login() echo.HandlerFunc
	Logout() echo.HandlerFunc
}

type authHandler struct {
	authUsecase usecase.IAuthUsecase
}

func NewHandler(authUsecase usecase.IAuthUsecase) IAuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
	}
}

func (h *authHandler) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.AuthRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		authUser, err := h.authUsecase.Signup(req.Username, req.Email, req.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := dto.AuthResponse{
			ID:       authUser.UserId,
			Username: authUser.UserId,
			Email:    authUser.Email,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

func (h *authHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.LoginRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		signedString, authUser, err := h.authUsecase.Login(req.Email, req.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		cookie := &http.Cookie{
			Name:   "jwt",
			Value:  signedString,
			MaxAge: 60 * 60 * 24,
			Path:   "/",
		}
		c.SetCookie(cookie)

		res := dto.LoginResponse{
			ID:       authUser.UserId,
			Username: authUser.UserId,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (h *authHandler) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie := &http.Cookie{
			Name:   "jwt",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successfully logged out",
		})
	}
}
