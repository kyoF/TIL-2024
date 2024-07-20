package main

import (
	"log"

	"backend/di"
	"backend/handler"
	"backend/infrastructure/mysql"
	"backend/infrastructure/redis"
)

func main() {
	mysqlDB, err := mysql.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	redisDB := redis.NewRedis()

	handler.InitRoute(di.InjectDependencies(mysqlDB, redisDB))
}
