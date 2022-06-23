package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

//check whether user is authorized or not
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			var err Error
			err = SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

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
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				r.Header.Set("Name", fmt.Sprintf("%v", claims["name"]))
				handler.ServeHTTP(w, r)
				return
			}
		}
		var reserr Error
		reserr = SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}
