package di

import (
	"gorm.io/gorm"

	"backend/application/usecase"
	mysqlinfra "backend/infrastructure/mysql"
	"backend/presentation/handler"
)

func New(db *gorm.DB) handler.User {
	return handler.NewUser(
		usecase.NewUser(
			mysqlinfra.NewUser(db),
		),
	)
}
