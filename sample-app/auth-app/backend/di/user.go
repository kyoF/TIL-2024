package di

import (
	"backend/application/usecase"
	"backend/handler/middleware"
	"backend/handler/router"
	mysqlPersistance "backend/infrastructure/mysql/persistance"
	redisPersistance "backend/infrastructure/redis/persistance"
	redisQueryService "backend/infrastructure/redis/queryservice"
	"database/sql"

	"github.com/go-redis/redis/v8"
)

func InjectDependencies(mysql *sql.DB, redis *redis.Client) (
	router.User,
	middleware.AuthN,
) {
	usecase := usecase.NewUser(
		mysqlPersistance.NewUser(mysql),
		redisPersistance.NewSession(redis),
		redisQueryService.NewSession(redis),
	)
	return router.NewUser(usecase), middleware.New(usecase)

}
