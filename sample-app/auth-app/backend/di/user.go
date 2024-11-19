package di

import (
	"backend/application/usecase"
	"backend/handler/middleware"
	"backend/handler/router"
	mysqlpersistance "backend/infrastructure/mysql/persistance"
	redispersistance "backend/infrastructure/redis/persistance"
	redisqueryservice "backend/infrastructure/redis/queryservice"
	"database/sql"

	"github.com/go-redis/redis/v8"
)

func InjectDependencies(mysql *sql.DB, redis *redis.Client) (
	router.User,
	middleware.AuthN,
) {
	usecase := usecase.NewUser(
		mysqlpersistance.NewUser(mysql),
		redispersistance.NewSession(redis),
		redisqueryservice.NewSession(redis),
	)
	return router.NewUser(usecase), middleware.New(usecase)

}
