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
		conn.updateRecord(rec)
	} else {
		conn.insertNewRecord()
	}
}

func updatePause(conn *connection, pause string) {
	rec := conn.getLastRecord()
	if rec.date == today() {
		conn.setPause(rec.date, pause)
	}
}

func main() {
	trackPtr := flag.Bool("track", false, "Track the current date and time")
	pausePtr := flag.String("pause", "", "Set the pause for today (use a time string like 30m or 1h10m)")
	statsPtr := flag.String("stats", "", "Print the statistics")
	flag.Parse()

	conn := connect("test.sqlite")
	conn.createTable()

	// ./tracker --track
	if *trackPtr == true {
		track(conn)
	}

	// ./tracker --pause {minutes}
	if *pausePtr != "" {
		updatePause(conn, *pausePtr)
	}

	// ./tracker --stats {timeSpan}
	if *statsPtr != "" {
		printStats(conn, *statsPtr)
	}
}
