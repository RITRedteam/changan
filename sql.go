package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func openDB(dsn string) *sql.DB {
	var db *sql.DB
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
