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
	defer mysqlDB.Close()

	redisDB := redis.NewRedis()
	defer redisDB.Close()

	handler.InitRoute(di.InjectDependencies(mysqlDB, redisDB))
}
