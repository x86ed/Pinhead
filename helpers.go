package main

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// load environment variable
func GoDotEnvVariable(key string) (string, error) {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		return "", err
	}

	env := os.Getenv(key)
	if len(env) < 1 {
		return "", errors.New("env var not found")
	}

	return env, nil
}

// set error message in Error struct
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

// take password as input and generate new hash password from it
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// compare plain password with hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Generate JWT token
func GenerateJWT(name, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetScoreState(s []Score, id uuid.UUID) string {
	for _, score := range s {
		if score.Active && score.User == id && score.Complete {
			return expired
		}
		if score.Active && score.User == id {
			return user
		}
	}
	return upcoming
}
