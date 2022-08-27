package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSignUp(t *testing.T) {

	data := url.Values{}
	data.Set("name", "ace")
	data.Set("initials", "ACE")
	data.Set("role", "user")

	req, err := http.NewRequest("POST", "/signup", strings.NewReader(data.Encode()))

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostSignUp)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("SignUp handler returned wrong status code")
	}
}

func TestSignIn(t *testing.T) {
	data := url.Values{}
	data.Set("name", "admin")
	data.Set("password", "test")

	req, err := http.NewRequest("POST", "/signin", strings.NewReader(data.Encode()))

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostSignIn)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("SignIn handler returned wrong status code")
	}
}

func TestUserIndex(t *testing.T) {
	data := url.Values{}
	data.Set("name", "ace")
	data.Set("role", "user")

	req, err := http.NewRequest("GET", "/game", strings.NewReader(data.Encode()))

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCurrentGame)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("UserIndex handler returned wrong status code")
	}
}

func TestListUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetListUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("ListUsers handler returned wrong status code")
	}
}

func TestDeleteAccount(t *testing.T) {
	data := url.Values{}
	data.Set("userId", "1")
	data.Set("name", "user")

	req, err := http.NewRequest("DELETE", "/admin/{userId}", strings.NewReader(data.Encode()))

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteAccount)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("ListUsers handler returned wrong status code")
	}
}
