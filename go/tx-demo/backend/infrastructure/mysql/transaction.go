package mysqlinfra

import "gorm.io/gorm"

type transaction struct {
	db *gorm.DB
	tx *gorm.DB
}

func (t *transaction) Transaction(f func() error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		t.tx = tx
		return f()
	})
}
