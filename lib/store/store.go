package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Pool struct {
	*sql.DB
}

var (
	DB *Pool
)

func Connect() error {
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		return err
	}
	DB = &Pool{db}
	return DB.Ping()
}
