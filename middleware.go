package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// using user jwt key
func IsAuthorizedUser(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authCook, _ := r.Cookie("Authorization")
		if len(authCook.Value) < 1 {
			authCook = nil
		}

		if r.Header["Authorization"] == nil && authCook == nil {
			var err Error
			err = SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		authVal := r.Header.Get("Authorization")
		if len(authVal) < 1 {
			authVal = authCook.Value
		}

		var mySigningKey = []byte(userSecretKey)
		token, err := jwt.Parse(strings.Replace(authVal, "Bearer ", "", 1), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err Error
			err = SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			handler.ServeHTTP(w, r)
			return
			// if the role claim is unknown then user will not be authorized
		}
		var reserr Error
		SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}

// using admin jwt key
func IsAuthorizedAdmin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			var err Error
			err = SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(adminSecretKey)
		token, err := jwt.Parse(strings.Replace(r.Header["Authorization"][0], "Bearer ", "", 1), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err Error
			err = SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r.Header.Set("Email", fmt.Sprintf("%v", claims["email"]))
			handler.ServeHTTP(w, r)
			return
		}
		var reserr Error
		SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}
