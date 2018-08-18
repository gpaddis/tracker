package main

import "time"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db := connect()
	createTable(db)

	id, recordDate := getLastRecord(db)
	currentDate := time.Now().Local().Format("2006-01-02")

	if currentDate != recordDate {
		insertNewRecord(db)
	} else {
		updateTodaysRecord(db, id)
	}
}
