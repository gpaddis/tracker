package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var dailyWorkHours, _ = time.ParseDuration("8h")

func today() string {
	return time.Now().Local().Format("2006-01-02")
}

func yesterday() string {
	return time.Now().Local().AddDate(0, 0, -1).Format("2006-01-02")
}

func printStats(timeSpan string, conn *connection) {
	switch timeSpan {
	case "today":
		printDailyReport(conn, today())
	case "yesterday":
		printDailyReport(conn, yesterday())
	case "thisweek":
		printWeeklyReport(conn, today())
	case "lastweek":
		fmt.Println("Last week")
	default:
		fmt.Println(timeSpan, "is not a valid time span.")
	}
}

func printDailyReport(conn *connection, date string) {
	rec := conn.getRecordByDay(date)
	if rec.date != "" {
		dur := calculateDuration(rec)
		fmt.Print(date+": ", dur.String(), "\t")

		balance := dur - dailyWorkHours
		fmt.Print("Balance: ")
		if balance < 0 {
			color.Set(color.FgRed)
		} else {
			color.Set(color.FgGreen)
			fmt.Print("+")
		}
		fmt.Println(balance.String())
		color.Unset()
	}
}

func printWeeklyReport(conn *connection, date string) {
	weekDays := getWeekDays(date)
	for _, day := range weekDays {
		printDailyReport(conn, day)
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

	return end.Sub(start) - pause
}

func getWeekDays(d string) []string {
	date, err := time.Parse("2006-01-02", d)
	checkErr(err)

	// Rewind to monday in the week
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
	}

	result := []string{}
	for len(result) < 7 {
		result = append(result, date.Format("2006-01-02"))
		date = date.AddDate(0, 0, 1)
	}

	return result
}
