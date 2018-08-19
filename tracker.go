package main

import (
	"flag"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func track(conn *connection) {
	rec := conn.getLastRecord()
	currentDate := time.Now().Local().Format("2006-01-02")

	if currentDate != rec.date {
		conn.insertNewRecord()
	} else {
		conn.updateRecord(rec.id)
	}
}

func main() {
	trackPtr := flag.Bool("track", false, "Track the current date and time")
	statsPtr := flag.String("stats", "", "Print the statistics")
	flag.Parse()

	conn := connect("test.sqlite")
	conn.createTable()

	// ./tracker --track
	if *trackPtr == true {
		track(conn)
	}

	// ./tracker --stats timeSpan
	if *statsPtr != "" {
		printStats(*statsPtr, conn)
	}
}
