package repository

import "app/src/domain/entity"

type IAuthRepository interface {
	CreateAuth(user entity.Auth) (*entity.Auth, error)
	GetAuthByEmail(email string) (*entity.Auth, error)
}
