package repository

import (
	"time"
)

type ISessionRepository interface {
	Insert(sessionId string, name string, expanded time.Duration) error
}
