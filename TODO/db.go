package main

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func initDB() {
	// connect to db
	var err error
	dbFilePath := "./tasks.db"
	DB, err = sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	defer DB.Close()

	// create table
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		notes TEXT,
		completed BOOLEAN,
		due_date TEXT,
	);`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Could not create table: %s\n", err.Error())
	}
}
