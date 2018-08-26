package main

import (
	"testing"
	"time"
)

func createTestRecord(date string) record {
	return record{1, date, "10:00:00", "18:00:00", "60m"}
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
	rec := createTestRecord("2018-08-20")
	workedHours, balance := getWorkedHours(&rec)
	expectedWorkedHours, _ := time.ParseDuration("7h")
	expectedBalance, _ := time.ParseDuration("-1h")
	if workedHours != expectedWorkedHours {
		t.Errorf("Expected 7 worked hours, got %s", workedHours)
	}
	if balance != expectedBalance {
		t.Errorf("Expected -1 balance hours, got %s", balance)
	}
}

func TestGetBalanceTotal(t *testing.T) {
	recordCollection := make([]record, 2)
	recordCollection[0] = createTestRecord(today())
	recordCollection[1] = createTestRecord(yesterday())

	balance := getBalanceTotal(recordCollection)
	expectedBalance := time.Duration(-2 * time.Hour)
	if balance != expectedBalance {
		t.Errorf("Expected a balance of %s, got %s", expectedBalance, balance)
	}
}
