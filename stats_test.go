package main

import (
	"testing"
	"time"
)

func createRecord() record {
	return record{1, "2018-08-20", "10:00:00", "18:00:00", "60m"}
}

func TestGetWeekDays(t *testing.T) {
	expected := []string{
		"2018-08-20",
		"2018-08-21",
		"2018-08-22",
		"2018-08-23",
		"2018-08-24",
		"2018-08-25",
		"2018-08-26",
	}

	actual := getWeekDays("2018-08-23")

	if len(expected) != len(actual) {
		t.Errorf("Expected 7 days, got %d", len(actual))
	}

	if expected[0] != actual[0] {
		t.Errorf("First date expected %s, got %s", expected[0], actual[0])
	}

	if expected[6] != actual[6] {
		t.Errorf("Last date expected %s, got %s", expected[6], actual[6])
	}
}

func TestGetWorkedHours(t *testing.T) {
	rec := createRecord()
	workedHours, balance := getWorkedHours(&rec)
	expectedWorkedHours, _ := time.ParseDuration("7h")
	expectedBalance, _ := time.ParseDuration("-1h")
	if workedHours != expectedWorkedHours {
		t.Errorf("Expected 7 worked hours, got %s", workedHours)
	}
	if balance != expectedBalance {
		t.Errorf("Expected 0 balance hours, got %s", balance)
	}
}
