package main

import (
	"log"

	"backend/application/usecase"
	"backend/handler"
	"backend/infrastructure/mysql"
	mysqlpersistance "backend/infrastructure/mysql/persistance"
	"backend/infrastructure/redis"
	redispersistance "backend/infrastructure/redis/persistance"
)

func main() {
	mysqlDB, err := mysql.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	redisDB := redis.NewRedis()

	handler.InitRoute(
		handler.NewUserHandler(
			usecase.NewUserUsecase(
				mysqlpersistance.NewMySQLPersistance(mysqlDB),
				redispersistance.NewRedisPersistance(redisDB),
			),
		),
	)
}
