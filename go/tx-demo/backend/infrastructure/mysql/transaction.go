package mysqlinfra

import "gorm.io/gorm"

type transaction struct {
	db *gorm.DB
}

func (t *transaction) Transaction(f func() error) error {
	var err error

	t.db = t.db.Begin()
	if t.db.Error != nil {
		return err
	}

	err = f()

	if err != nil {
		t.db = t.db.Rollback()
		if t.db.Error != nil {
			return err
		}
	}

	t.db = t.db.Commit()
	if t.db.Error != nil {
		t.db = t.db.Rollback()
		if t.db.Error != nil {
			return err
		}
	}

	t.db, err = NewDB()
	if err != nil {
		return err
	}

	return nil
}
