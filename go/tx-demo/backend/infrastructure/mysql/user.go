package mysqlinfra

import "backend/domain/repository"

type user struct {
    *transaction
}

func NewUser(*transaction transaction) repository.User {
    
}
