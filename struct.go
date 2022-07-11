package main

import uuid "github.com/satori/go.uuid"

type Authentication struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Name        string `json:"name"`
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
	Name     string
	Initials string
	Class    string
}

type ScoreUpdate struct {
	ID    uuid.UUID
	Score int
}
