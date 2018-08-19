package main

import (
	"fmt"
	"time"
)

func today() string {
	return time.Now().Local().Format("2006-01-02")
}

func yesterday() string {
	return time.Now().Local().AddDate(0, 0, -1).Format("2006-01-02")
}

func printStats(timeSpan string, conn *connection) {
	switch timeSpan {
	case "today":
		getDailyReport(conn, today())
	case "yesterday":
		getDailyReport(conn, yesterday())
	case "thisweek":
		fmt.Println("This week")
	case "lastweek":
		fmt.Println("Last week")
	default:
		fmt.Println(timeSpan, "is not a valid time span.")
	}
}

func getDailyReport(conn *connection, date string) {
	rec := conn.getRecordByDay(date)
	if rec.date != "" {
		dur := calculateDuration(rec)
		fmt.Println("Stats for "+date+":", dur.String())
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
	if r.date != today() {
		delta -= pause
	}
	return delta
}
