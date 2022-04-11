package db

import (
	"database/sql"
	"fmt"
)

func SqlConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:Prado393@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		fmt.Println(err)
	}

	return db
}
