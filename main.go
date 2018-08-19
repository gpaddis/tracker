package main

import "time"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn := connect("test.sqlite")
	conn.createTable()

	id, recordDate := conn.getLastRecord()
	currentDate := time.Now().Local().Format("2006-01-02")

	if currentDate != recordDate {
		conn.insertNewRecord()
	} else {
		conn.updateRecord(id)
	}
}
