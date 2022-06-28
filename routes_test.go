package main

import "testing"

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCreateRouter(t *testing.T) {
	CreateRouter()
	if router == nil {
		t.Error("Create router failed")
	}
}

func TestInitializeStatic(t *testing.T) {
	InitializeStatic()
	if router == nil {
		t.Error("Create router failed")
	}
}

func TestInitializeRoute(t *testing.T) {
	InitializeRoute()
	if router == nil {
		t.Error("Create router failed")
	}
}
