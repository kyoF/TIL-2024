package presenter

import (
	"context"

	"RestAPI/api/internal/controller/system"
	"RestAPI/api/internal/controller/user"

	"github.com/gin-gonic/gin"
)

const latest = "/v1"

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)

	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	{
		userHandler := user.NewUserHandler()
		v1.GET("", userHandler.GetUser)
		v1.GET("/:id", userHandler.GetUserById)
		v1.POST("", userHandler.EditUser)
		v1.DELETE("", userHandler.DeleteUser)
	}

	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}
