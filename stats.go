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

func oneWeekAgo() string {
	return time.Now().Local().AddDate(0, 0, -7).Format("2006-01-02")
}

func printStats(conn *connection, timeSpan string) {
	switch timeSpan {
	case "today":
		workedHours, balance := getWorkedHoursByDay(conn, today())
		printDailyReport(today(), workedHours, balance)
	case "yesterday":
		workedHours, balance := getWorkedHoursByDay(conn, today())
		printDailyReport(yesterday(), workedHours, balance)
	case "thisweek":
		printWeeklyReport(conn, today())
	case "lastweek":
		printWeeklyReport(conn, oneWeekAgo())
	default:
		fmt.Println(timeSpan, "is not a valid time span.")
	}
}

func getWorkedHours(rec *record) (time.Duration, time.Duration) {
	var workedHours, balance time.Duration
	if rec.date != "" {
		workedHours = calculateDuration(rec)
		balance = workedHours - dailyWorkHours
	}
	return workedHours, balance
}

func getWorkedHoursByDay(conn *connection, date string) (time.Duration, time.Duration) {
	rec := conn.getRecordByDay(date)
	return getWorkedHours(rec)
}

func printDailyReport(date string, workedHours time.Duration, balance time.Duration) {
	if workedHours != 0 {
		fmt.Print(date+": ", workedHours.String(), "\t")
		printBalance(balance)
	}
}

func printBalance(b time.Duration) {
	fmt.Print("Balance: ")
	if b < 0 {
		color.Set(color.FgRed)
	} else {
		color.Set(color.FgGreen)
		fmt.Print("+")
	}
	fmt.Println(b.String())
	color.Unset()
}

func printWeeklyReport(conn *connection, date string) {
	weekDays := getWeekDays(date)
	var finalBalance time.Duration
	for _, day := range weekDays {
		workedHours, balance := getWorkedHoursByDay(conn, day)
		printDailyReport(day, workedHours, balance)
		finalBalance += balance
	}
	fmt.Print("\t\t\tWeekly ")
	printBalance(finalBalance)
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
