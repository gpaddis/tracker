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

func updatePause(conn *connection, pause int) {
	rec := conn.getLastRecord()
	if rec.date == today() {
		conn.setPause(rec.date, pause)
	}
}

func main() {
	trackPtr := flag.Bool("track", false, "Track the current date and time")
	pausePtr := flag.Int("pause", 0, "Set the pause for today (in minutes)")
	statsPtr := flag.String("stats", "", "Print the statistics")
	flag.Parse()

	conn := connect("test.sqlite")
	conn.createTable()

	// ./tracker --track
	if *trackPtr == true {
		track(conn)
	}

	// ./tracker --pause {minutes}
	if *pausePtr != 0 {
		updatePause(conn, *pausePtr)
	}

	// ./tracker --stats {timeSpan}
	if *statsPtr != "" {
		printStats(*statsPtr, conn)
	}
}
