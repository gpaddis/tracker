package main

import (
	"fmt"
)

func printStats(timeSpan string) {
	switch timeSpan {
	case "today":
		fmt.Println("Today")
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
