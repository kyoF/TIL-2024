package mysqlinfra

import (
	"backend/domain/repository"

	"gorm.io/gorm"
)

type transaction struct {
	db *gorm.DB
}

func (t *transaction) Exec(f func(tx *transaction)) error {
	return t.db.Transaction(f(tx))
}
