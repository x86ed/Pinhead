package main

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(db *gorm.DB) error {
	base.ID = uuid.NewV4()
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return nil
}

type User struct {
	Base
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Initials string `json:"initials"`
}

type Score struct {
	Base
	User     uuid.UUID `json:"user"`
	Score    int64     `json:"score"`
	Complete bool      `json:"complete"`
	Active   bool      `json:"active"`
}

type Game struct {
	Base
	Scores []Score `json:"scores" gorm:"many2many:GameScore"`
	Active bool    `json:"active"`
	Users  []User  `json:"users" gorm:"many2many:GameUser"`
}
