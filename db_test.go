package main

import "testing"

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGetDatabase(t *testing.T) {
	val, err := GetDatabase()
	if val == nil {
		t.Error("Database not Returned")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestInitalMigration(t *testing.T) {
	err := InitialMigration()
	if err != nil {
		t.Error(err)
	}
}
