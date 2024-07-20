package repository

import "time"

type ISessionRepository interface {
	Set(sessionId, name string, deadline time.Duration) error
	Delete(sessionId string) error
}
