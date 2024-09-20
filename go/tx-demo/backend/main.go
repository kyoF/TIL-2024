package main

import (
	"log"

	"backend/di"

	mysqlinfra "backend/infrastructure/mysql"
	"backend/infrastructure/mysql/models"
	"backend/presentation"
)

func main() {
	db, err := mysqlinfra.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Item{})

	presentation.InitRoute(di.New(db))
}
