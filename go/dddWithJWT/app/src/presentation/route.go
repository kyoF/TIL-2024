package presentation

import (
	"dddWithJWT/pkg/infrastructure"
	"dddWithJWT/pkg/infrastructure/repositoryimpl"
	"dddWithJWT/pkg/interfaces/api/handler"
	"dddWithJWT/pkg/interfaces/api/middleware"
	"dddWithJWT/pkg/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func InitRoute() {
	userRepo := repositoryimpl.NewRepositoryImpl(infrastructure.Conn)
	userUseCase := usecase.NewUseCase(userRepoImpl)
	userHandler := handler.NewHandler(userUseCase)

	r := echo.New()

	r.POST("/signup", userHandler.HandleSignup)
	r.POST("/login", userHandler.HandleLogin)
	r.GET("/logout", userHandler.HandleLogout)

	withAuth := r.Group("/auth").Use(middleware.Auth())
	withAuth.GET("/ping", Ping)

	log.Println("Server running...")
	if err := r.Run(); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
