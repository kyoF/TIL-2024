package user

import "github.com/gin-gonic/gin"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUser(ctx *gin.Context) {}

func (h *UserHandler) GetUserById(ctx *gin.Context) {}

func (h *UserHandler) EditUser(ctx *gin.Context) {}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {}
