package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	return sql.Open(
		"mysql",
		"sqlboiler:sqlboiler@tcp(database:3306)/sqlboiler?parseTime=true",
	)
}
