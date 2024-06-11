package dao

import "database/sql"

func NewDB() (*sql.DB, error) {
	return sql.Open(
		"mysql",
		"host=sqlboiler dbname=sqlboiler user=sqlboiler password=sqlboiler sslmode=disable",
	)
}
