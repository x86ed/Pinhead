package main

import (
	"testing"
)

func TestGoDotEnvVariable(t *testing.T) {
	o, err := GoDotEnvVariable("JWTKEY")
	if err != nil {
		t.Error("JWTKEY not found")
	}
	if o != "secretkeyjwt" {
		t.Error("wrong value")
	}
	_, err = GoDotEnvVariable("FAKE")
	if err == nil {
		t.Error("Fake env var shouldn't have value")
	}
}

func TestSetError(t *testing.T) {
	var e Error
	ee := SetError(e, "test")
	if !ee.IsError && ee.Message != "test" {
		t.Error("error not set correctly")
	}
}

func TestGeneratehashPassword(t *testing.T) {
	s, err := GenerateHashPassword("test")
	if err != nil {
		t.Error("bad hash")
	}
	if len(s) < 1 {
		t.Error("bad hash")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	s, _ := GenerateHashPassword("test")
	if !CheckPasswordHash("test", s) {
		t.Error("bad hash")
	}
}

func TestGenerateJWT(t *testing.T) {
	s, err := GenerateJWT("test", "user")
	if err != nil {
		t.Error("JWT failed")
	}
	if len(s) < 1 {
		t.Error("JWT failed")
	}
}
