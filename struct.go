package main

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// // BeforeCreate will set a UUID rather than numeric ID.
// func (base *Base) BeforeCreate(db *gorm.DB) error {
// 	uuid := uuid.NewV4()
// 	return db.Set("ID", uuid)
// }

type User struct {
	Base
	Name     string `gorm:"unique" json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Initials string `json:"initials"`
}

type Score struct {
	Base
	User  uuid.UUID `json:"user"`
	Score int64     `json:"score"`
}

type Game struct {
	Base
	Scores  []Score
	Active  bool `json:"active"`
	Players []User
}

type Authentication struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
