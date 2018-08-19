package main

// See:
// * http://www.sqlitetutorial.net/sqlite-date/
// * Embedding a pointer in a struct: https://stackoverflow.com/a/44412899

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type connection struct {
	*sql.DB
}

type record struct {
	id        int
	date      string
	startTime string
	endTime   string
}

func connect(dbName string) *connection {
	conn, err := sql.Open("sqlite3", dbName)
	checkErr(err)
	return &connection{conn}
}

func (c connection) executeStmt(statement string) sql.Result {
	stmt, err := c.Prepare(statement)
	checkErr(err)
	res, err := stmt.Exec()
	return res
}

func (c connection) createTable() sql.Result {
	stmt := "CREATE TABLE IF NOT EXISTS tracker (id INTEGER PRIMARY KEY, date TEXT, start_time TEXT, end_time TEXT)"
	return c.executeStmt(stmt)
}

func (c connection) insertNewRecord() sql.Result {
	return c.executeStmt("INSERT INTO tracker (date, start_time) VALUES (DATE('now', 'localtime'), TIME('now', 'localtime'))")
}

func (c connection) updateRecord(id int) sql.Result {
	stmt, err := c.Prepare("UPDATE tracker SET end_time = TIME('now', 'localtime') WHERE id = ?")
	checkErr(err)
	res, err := stmt.Exec(id)
	return res
}

func (c connection) getLastRecord() (int, string) {
	row, err := c.Query("SELECT id, date FROM tracker ORDER BY id DESC LIMIT 1")
	checkErr(err)

	var id int
	var date string

	for row.Next() {
		err := row.Scan(&id, &date)
		checkErr(err)
	}
	return id, date
}

func (c connection) getRecordByDay(day string) *record {
	row, err := c.Query("SELECT * FROM tracker WHERE date = ? LIMIT 1", day)
	checkErr(err)
	r := new(record)
	for row.Next() {
		err := row.Scan(&r.id, &r.date, &r.startTime, &r.endTime)
		checkErr(err)
	}
	return r
}
