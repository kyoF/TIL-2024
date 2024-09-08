package mysqlinfra

import "gorm.io/gorm"

type transaction struct {
	db   *gorm.DB
	txDB *gorm.DB
}

func (t *transaction) Transaction(f func() error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		t.txDB = tx
		return f()
	})
}
