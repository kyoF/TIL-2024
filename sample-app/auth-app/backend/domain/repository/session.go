package repository

import "time"

type Session interface {
	Set(sessionId, name string, deadline time.Duration) error
	Delete(sessionId string) error
}
