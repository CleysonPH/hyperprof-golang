package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func InitMySQL(dsn string) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func Close() {
	db.Close()
}
