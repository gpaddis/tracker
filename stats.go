package main

import (
	"fmt"
	"time"
)

func today() string {
	return time.Now().Local().Format("2006-01-02")
}

func printStats(timeSpan string, conn *connection) {
	switch timeSpan {
	case "today":
		rec := conn.getRecordByDay(today())
		dur := calculateDuration(rec)
		fmt.Println("Stats for today:", dur.String())
	case "yesterday":
		fmt.Println("Yesterday")
	case "thisweek":
		fmt.Println("This week")
	case "lastweek":
		fmt.Println("Last week")
	default:
		fmt.Println(timeSpan, "is not a valid time span.")
	}
}

func calculateDuration(r *record) time.Duration {
	dateFormat := "2006-01-0215:04:05"

	start, err := time.Parse(dateFormat, r.date+r.startTime)
	checkErr(err)

	end, err := time.Parse(dateFormat, r.date+r.endTime)
	checkErr(err)

	pause, err := time.ParseDuration(r.pause)
	checkErr(err)

	delta := end.Sub(start)
	if delta > pause {
		delta -= pause
	}
	return delta
}
