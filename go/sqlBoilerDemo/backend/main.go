package main

import (
	"backend/application"
	"backend/infrastructure"
	"backend/infrastructure/dao"
	"backend/presentation"
)

func main() {
	db, _ := dao.NewDB()

	presentation.InitRoute(
		presentation.NewRouter(
			application.NewUsecase(
				infrastructure.NewInfrastructure(db),
			),
		),
	)
}
