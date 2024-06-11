package infrastructure

import (
	"backend/domain"
	"database/sql"
)

type mysqlInfrastructure struct {
	db *sql.DB
}

func NewInfrastructure(db *sql.DB) domain.IRepository {
	return &mysqlInfrastructure{
		db: db,
	}
}

func (i *mysqlInfrastructure) GetUsers() {}
