package main

import (
	"backend/di"
	mysqlinfra "backend/infrastructure/mysql"
	handler "backend/presentation"
	"log"
)

func main() {
	mysqlDB, err := mysqlinfra.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoute(di.New(mysqlDB))
}
