package main

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
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
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Initials string `json:"initials" gorm:"type:varchar(3)"`
}

type Admin struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password"`
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
	Scores   []Score `json:"scores" gorm:"many2many:GameScore"`
	InActive bool    `json:"in_active"`
	Users    []User  `json:"users" gorm:"many2many:GameUser"`
}

type Control struct {
	Base
	DomID       string         `json:"dom_id"`
	Keys        datatypes.JSON `json:"keys"`
	DownCommand string         `json:"down_command"`
	UpCommand   string         `json:"up_command"`
}
