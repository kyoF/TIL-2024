package repository

import "backend/domain/entity"

type User interface {
	Get(userId string) (*entity.User, error)
	Insert(userId, name string, age int) error
	UpdateName(userId, name string) error
	UpdateAge(userId string, age int) error
	DBClient
}
