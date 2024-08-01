package main

import (
	"backend/application"
	"backend/infrastructure"
	"backend/infrastructure/dao"
	"backend/presentation"
	"log"
)

func main() {
	db, err := dao.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	presentation.InitRoute(
		presentation.NewRouter(
			application.NewUsecase(
				infrastructure.NewInfrastructure(db),
			),
		),
	)
}
