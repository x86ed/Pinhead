package main

import uuid "github.com/satori/go.uuid"

type Authentication struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AdminAuthentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	TokenString string `json:"token"`
}

type AdminToken struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Error struct {
	IsError bool   `json:"is_error"`
	Message string `json:"message"`
}

const (
	user     = "user"
	expired  = "expired"
	upcoming = "upcoming"
)

type Player struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Initials string    `json:"initials"`
	Class    string    `json:"class"`
	Score    int       `json:"score"`
}

type CurGame struct {
	Players []Player `json:"players"`
	CurID   string   `json:"cur_id"`
}

type ScoreUpdate struct {
	ID    string `json:"id"`
	Score string `json:"score"`
}
