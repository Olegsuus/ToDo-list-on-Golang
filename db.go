package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateTable(db *sql.DB) {
	createTasksTableSql := `CREATE TABLE IF NOT EXISTS tasks (
		"idTask" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"category_id" INTEGER,
		"completed" INTEGER DEFAULT 0,
		"createdAt" DATETIME DEFAULT CURRENT_TIMESTAMP,
		"updatedAt" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(createTasksTableSql)
	if err != nil {
		log.Fatal(err)
	}
}
