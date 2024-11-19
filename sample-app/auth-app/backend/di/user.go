package di

import (
	"backend/application/usecase"
	"backend/handler/middleware"
	"backend/handler/router"
	mysqlpersistance "backend/infrastructure/mysql/persistance"
	queryserviceimpl "backend/infrastructure/redis/queryservice"
	redispersistance "backend/infrastructure/redis/persistance"
	"database/sql"

	"github.com/go-redis/redis/v8"
)

func InjectDependencies(mysql *sql.DB, redis *redis.Client) (
	router.IUserRouter,
	middleware.IMiddleware,
) {
	usecase := usecase.NewUserUsecase(
		mysqlpersistance.NewMySQLPersistance(mysql),
		redispersistance.NewRedisPersistance(redis),
		queryserviceimpl.NewRedisQueryService(redis),
	)
	return router.NewUserRouter(usecase), middleware.NewMiddleware(usecase)

}
