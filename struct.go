package main

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
	Name     string `json:"name"`
	Initials string `json:"initials"`
	Class    string `json:"class"`
	Score    int64  `json:"score"`
}

type AdminPlayer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Initials string `json:"initials"`
	Class    string `json:"class"`
	Score    int64  `json:"score"`
}

type CurGame struct {
	Players []Player `json:"players"`
	CurID   string   `json:"cur_id"`
}

type ScoreUpdate struct {
	ID    string
	Score int
}

type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
