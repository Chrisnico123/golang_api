package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql" , "root@tcp(localhost:3306)/db_golang")
	if err != nil {
		panic(err)
	} 

	return db
}