package repository

import "time"

type IAuthRepository interface {
	Set(sessionId, name string, deadline time.Duration) error
	Delete(sessionId string) error
	Get(sessionId string) (string, error)
}
