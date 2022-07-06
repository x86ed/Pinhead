package main

import (
	"testing"

	"gorm.io/gorm"
)

func TestBeforeCreate(t *testing.T) {
	db := &gorm.DB{}
	b := Base{}
	bb := Base{}
	err := b.BeforeCreate(db)
	if err != nil {
		t.Error("before create failed")
	}
	if b.ID == bb.ID {
		t.Error("bad uuid")
	}
}
