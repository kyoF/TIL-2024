package repository

import "backend/domain/entity"

type IUserRepository interface {
	Insert(name, password string) error
	Get(name string) (entity.User, error)
}
