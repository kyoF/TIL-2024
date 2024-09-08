package main

import (
	"log"

	"backend/di"

	mysqlinfra "backend/infrastructure/mysql"
	"backend/infrastructure/mysql/models"
	"backend/presentation"
)

func main() {
	mysqlDB, err := mysqlinfra.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	mysqlDB.AutoMigrate(&models.User{})

	presentation.InitRoute(di.New(mysqlDB))
}
