package di

import (
	"gorm.io/gorm"

	"backend/application/usecase"
	mysqlinfra "backend/infrastructure/mysql"
	"backend/presentation/handler"
)

func New(db *gorm.DB) (
	handler.User,
	handler.Item,
	handler.Sample,
) {
	userRepository := mysqlinfra.NewUser(db)
	userUsecase := usecase.NewUser(userRepository)
	userHandler := handler.NewUser(userUsecase)

	itemRepository := mysqlinfra.NewItem(db)
	itemUsecase := usecase.NewItem(itemRepository)
	itemHandler := handler.NewItem(itemUsecase)

	sampleUsecase := usecase.NewSample(userRepository, itemRepository)
	sampleHandler := handler.NewSample(sampleUsecase)

	return userHandler, itemHandler, sampleHandler
}
