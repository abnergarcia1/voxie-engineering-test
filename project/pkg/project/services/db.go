package services

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func getDB() (db *sql.DB, e error) {
	user := "contacts"
	pass := "secret"
	host := "tcp(localhost:3306)"
	dbName := "contacts"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true", user, pass, host, dbName))
	if err != nil {
		return nil, err
	}
	return db, nil
}
