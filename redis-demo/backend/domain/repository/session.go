package repository

import (
	"time"
)

type ISessionRepository interface {
	Set(sessionId string, name string, expanded time.Duration) error
    Delete(sessionId string) error
}
