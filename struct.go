package main

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
