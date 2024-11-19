package redispersistance

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"backend/domain/repository"
)

type infra struct {
	db *redis.Client
}

func NewRedisPersistance(db *redis.Client) repository.IAuthRepository {
	return &infra{db: db}
}

func (i *infra) Set(sessionId, name string, deadline time.Duration) error {
	ctx := context.Background()
	err := i.db.Set(ctx, sessionId, name, deadline).Err()
	return err
}

func (i *infra) Delete(sessionId string) error {
	ctx := context.Background()
	err := i.db.Del(ctx, sessionId).Err()
	return err
}
