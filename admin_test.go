package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestNewGame(t *testing.T) {

	req, err := http.NewRequest("POST", "/new_game", nil)

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostNewGame)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("NewGame admin handler returned wrong status code")
	}
}

func TestNextTurn(t *testing.T) {

	req, err := http.NewRequest("POST", "/next_turn", nil)

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostNextTurn)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("NextTurn admin handler returned wrong status code")
	}
}

func TestUpdateScore(t *testing.T) {

	data := url.Values{}
	data.Set("id", "1")
	data.Set("score", "1")

	req, err := http.NewRequest("POST", "/update_score", strings.NewReader(data.Encode()))

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostUpdateScore)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("UpdateScore admin handler returned wrong status code")
	}
}

func TestHighScore(t *testing.T) {

	data := url.Values{}
	data.Set("id", "1")
	data.Set("score", "1")

	req, err := http.NewRequest("POST", "/high_score", strings.NewReader(data.Encode()))

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostHighScore)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("HighScore admin handler returned wrong status code")
	}
}

func TestCreateAdmin(t *testing.T) {

	data := url.Values{}
	data.Set("name", "admin")
	data.Set("role", "admin")

	req, err := http.NewRequest("POST", "/admin", strings.NewReader(data.Encode()))

	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostCreateAdmin)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("CreateAdmin admin handler returned wrong status code")
	}
}
