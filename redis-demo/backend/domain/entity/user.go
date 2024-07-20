package entity

import "backend/domain/valueobject"

type User struct {
	Name     string
	Password valueobject.Password
}
