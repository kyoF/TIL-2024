package repository

import "dddWithJWT/domain/entity"

type AuthRepository interface {
	CreateAuth(user entity.Auth) (*entity.Auth, error)
	GetAuthByEmail(email string) (*entity.Auth, error)
}
