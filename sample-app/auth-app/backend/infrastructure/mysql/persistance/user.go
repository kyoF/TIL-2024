package mysqlpersistance

import (
	"backend/domain/repository"
	"database/sql"
)

type infra struct {
	db *sql.DB
}

func NewUser(db *sql.DB) repository.User {
	return &infra{db: db}
}

func (i *infra) Insert(name, password string) error {
	_, err := i.db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", name, password)
	return err
}

func (i *infra) Get(name string) (string, error) {
	var password string
	err := i.db.QueryRow("SELECT password FROM users WHERE name = ?", name).Scan(&password)
	return password, err
}
