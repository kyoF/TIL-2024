package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "mysqluser:mysqlpassword@tcp(mysql:3306)/redisdemo")
	return db, err
}
