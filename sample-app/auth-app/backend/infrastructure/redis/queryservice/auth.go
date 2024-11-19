package queryservice

import (
	queryservice "backend/application/queryservice/interface"
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type infra struct {
	db *redis.Client
}

func NewRedisQueryService(db *redis.Client) queryservice.Auth {
	return &infra{db: db}
}

func (i *infra) Get(sessionId string) (string, error) {
	ctx := context.Background()
	name, err := i.db.Get(ctx, sessionId).Result()
	if err == redis.Nil {
		return "", errors.New("Invalid session ID")
	}
	return name, err
}
