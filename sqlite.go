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
	pause     int
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
	stmt := "CREATE TABLE IF NOT EXISTS tracker (id INTEGER PRIMARY KEY, date TEXT, start_time TEXT, end_time TEXT, pause INTEGER)"
	return c.executeStmt(stmt)
}

func (c connection) insertNewRecord() sql.Result {
	return c.executeStmt("INSERT INTO tracker (date, start_time, end_time, pause) VALUES (DATE('now', 'localtime'), TIME('now', 'localtime'), TIME('now', 'localtime'), 60)")
}

func (c connection) updateRecord(id int) sql.Result {
	stmt, err := c.Prepare("UPDATE tracker SET end_time = TIME('now', 'localtime') WHERE id = ?")
	checkErr(err)
	res, err := stmt.Exec(id)
	checkErr(err)
	return res
}

func (c *connection) getLastRecord() *record {
	row, err := c.Query("SELECT * FROM tracker ORDER BY id DESC LIMIT 1")
	checkErr(err)
	r := new(record)
	for row.Next() {
		err := row.Scan(&r.id, &r.date, &r.startTime, &r.endTime, &r.pause)
		checkErr(err)
	}
	return r
}

func (c connection) getRecordByDay(day string) *record {
	row, err := c.Query("SELECT * FROM tracker WHERE date = ? LIMIT 1", day)
	checkErr(err)
	r := new(record)
	for row.Next() {
		err := row.Scan(&r.id, &r.date, &r.startTime, &r.endTime, &r.pause)
		checkErr(err)
	}
	return r
}

func (c connection) setPause(day string, pause int) sql.Result {
	stmt, err := c.Prepare("UPDATE tracker SET pause = ? WHERE date = ?")
	checkErr(err)
	res, err := stmt.Exec(pause, day)
	checkErr(err)
	return res
}
