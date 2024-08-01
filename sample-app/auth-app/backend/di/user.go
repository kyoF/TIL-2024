package di

import (
	"database/sql"

	"github.com/go-redis/redis/v8"

	"backend/application/queryservice"
	"backend/application/usecase"
	"backend/handler/middleware"
	"backend/handler/router"
	mysqlpersistance "backend/infrastructure/mysql/persistance"
	redispersistance "backend/infrastructure/redis/persistance"
)

func InjectDependencies(mysql *sql.DB, redis *redis.Client) (
	router.IUserRouter,
	middleware.IMiddleware,
) {
	return router.NewUserRouter(
			usecase.NewUserUsecase(
				mysqlpersistance.NewMySQLPersistance(mysql),
				redispersistance.NewRedisPersistance(redis),
			),
		),
		middleware.NewMiddleware(
			queryservice.NewAuthQueryService(
				redispersistance.NewRedisPersistance(redis),
			),
		)

}
