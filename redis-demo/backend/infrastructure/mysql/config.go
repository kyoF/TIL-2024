package mysql

import "database/sql"

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "mysqluser:mysqlpassword@tcp(mysql:3306)/redisdemo")
	return db, err
}
