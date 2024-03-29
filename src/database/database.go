package database

import (
	"database/sql"
	"goapi/src/config"

	_ "github.com/go-sql-driver/mysql"
)

// this func opens connection with the db and returns it
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionStringDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
