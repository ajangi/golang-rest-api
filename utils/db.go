package utils

import (
	"database/sql" // Database SQL package to perform queries
	// Display messages to console
	// Manage URL
	// Manage HTML files
	_ "github.com/go-sql-driver/mysql" // MySQL Database driver
)

func DbConn() (db *sql.DB) {
	// Realize the connection with mysql driver
	db, err := sql.Open("mysql", "admin:megadeth1122@tcp(localhost:8889)/nardoon_db?charset=utf8&parseTime=True&loc=Local")

	// If error stop the application
	if err != nil {
		panic(err.Error())
	}
	// Return db object to be used by other functions
	return db
}
