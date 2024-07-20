package di

import (
	"database/sql"

	"github.com/go-redis/redis/v8"

	"backend/application/usecase"
	"backend/handler"
	mysqlpersistance "backend/infrastructure/mysql/persistance"
	redispersistance "backend/infrastructure/redis/persistance"
)

func InjectDependencies(mysql *sql.DB, redis *redis.Client) handler.IUserHandler {
	return handler.NewUserHandler(
		usecase.NewUserUsecase(
			mysqlpersistance.NewMySQLPersistance(mysql),
			redispersistance.NewRedisPersistance(redis),
		),
	)
}
