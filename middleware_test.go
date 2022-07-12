package main

import (
	"net/http"
	"testing"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, User."))
}

func TestIsAuthorized(t *testing.T) {
	IsAuthorized(testHandler, false)
}
