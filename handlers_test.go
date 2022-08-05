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
	handler := http.HandlerFunc(SignUp)
	
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Error("signup handler returned wrong status code")
	}
	
}