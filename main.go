package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *sql.DB {
	db, err := sql.Open("sqlite3", "test.sqlite")
	checkErr(err)
	return db
}

func createTable(db *sql.DB) sql.Result {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS tracker (id INTEGER PRIMARY KEY, date TEXT, start_time TEXT, end_time TEXT)")
	checkErr(err)
	res, err := stmt.Exec()
	return res
}

func insertNewRecord(db *sql.DB) sql.Result {
	// http://www.sqlitetutorial.net/sqlite-date/
	stmt, err := db.Prepare("INSERT INTO tracker (date, start_time) VALUES (DATE('now', 'localtime'), TIME('now', 'localtime'))")
	checkErr(err)
	res, err := stmt.Exec()
	return res
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

func main() {
	db := connect()
	createTable(db)

	id, date := getLastRecord(db)
	currentDate := time.Now().Local().Format("2006-01-02")
	if currentDate != date {
		insertNewRecord(db)
	} else {
		updateTodaysRecord(db, id)
	}
}
