package main

import (
	"testing"
)

func connectToTestDb() *connection {
	c := connect(":memory:")
	c.createTable()
	return c
}

func TestInsertNewRecord(t *testing.T) {
	c := connectToTestDb()
	res := c.insertNewRecord()
	id, _ := res.LastInsertId()
	if id != 1 {
		t.Error("Something went wrong when inserting a new record.")
	}
}

func TestUpdateRecord(t *testing.T) {
	c := connectToTestDb()
	c.insertNewRecord()
	rec := c.getLastRecord()
	res := c.updateRecord(rec)
	id, _ := res.LastInsertId()
	if id != 1 {
		t.Error("Something went wrong when updating a record.")
	}
}

func TestUpdateRecordWithEmptyDatabase(t *testing.T) {
	c := connectToTestDb()
	rec := c.getLastRecord()
	res := c.updateRecord(rec)
	id, _ := res.LastInsertId()
	if id != 0 {
		t.Error("Did not expect to get id when no records are present in the database.")
	}
}
