package mysqlinfra

import "gorm.io/gorm"

type dbClient struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewDBClient(db *gorm.DB) *dbClient {
	return &dbClient{db: db}
}

func (t *dbClient) Transaction(f func() error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		t.tx = tx
		return f()
	})
}

func (t *dbClient) DB() *gorm.DB {
	if t.tx == nil {
		return t.db
	}
	return t.tx
}
