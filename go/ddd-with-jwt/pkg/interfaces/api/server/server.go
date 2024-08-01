package server

import (
	"dddWithJWT/pkg/infrastructure"
	"dddWithJWT/pkg/infrastructure/repositoryimpl"
	"dddWithJWT/pkg/interfaces/api/handler"
	"dddWithJWT/pkg/interfaces/api/middleware"
	"dddWithJWT/pkg/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func Serve(addr string) {
	userRepoImpl := repositoryimpl.NewRepositoryImpl(infrastructure.Conn)
	userUseCase := usecase.NewUseCase(userRepoImpl)
	userHandler := handler.NewHandler(userUseCase)

	r = gin.Default()

	r.POST("/signup", userHandler.HandleSignup)
	r.POST("/login", userHandler.HandleLogin)
	r.GET("/logout", userHandler.HandleLogout)

	secured := r.Group("/secured").Use(middleware.Auth())
	secured.GET("/ping", Ping)

	log.Println("Server running...")
	if err := r.Run(addr); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
