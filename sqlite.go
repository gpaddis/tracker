package main

// See: http://www.sqlitetutorial.net/sqlite-date/

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func connect() *sql.DB {
	db, err := sql.Open("sqlite3", "test.sqlite")
	checkErr(err)
	return db
}

func executeStmt(db *sql.DB, statement string) sql.Result {
	stmt, err := db.Prepare(statement)
	checkErr(err)
	res, err := stmt.Exec()
	return res
}

func createTable(db *sql.DB) sql.Result {
	return executeStmt(db, "CREATE TABLE IF NOT EXISTS tracker (id INTEGER PRIMARY KEY, date TEXT, start_time TEXT, end_time TEXT)")
}

func insertNewRecord(db *sql.DB) sql.Result {
	return executeStmt(db, "INSERT INTO tracker (date, start_time) VALUES (DATE('now', 'localtime'), TIME('now', 'localtime'))")
}

func updateTodaysRecord(db *sql.DB, id int) sql.Result {
	stmt, err := db.Prepare("UPDATE tracker SET end_time = TIME('now', 'localtime') WHERE id = ?")
	checkErr(err)
	res, err := stmt.Exec(id)
	return res
}

func getLastRecord(db *sql.DB) (int, string) {
	row, err := db.Query("SELECT id, date FROM tracker ORDER BY id DESC LIMIT 1")
	checkErr(err)

	var id int
	var date string

	for row.Next() {
		err := row.Scan(&id, &date)
		checkErr(err)
	}
	return id, date
}
