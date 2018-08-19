package main

import (
	"flag"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func track(conn *connection) {
	rec := conn.getLastRecord()
	if rec.date == today() {
		conn.updateRecord(rec.id)
	} else {
		conn.insertNewRecord()
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
