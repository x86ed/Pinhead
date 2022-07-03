package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var dbuser User
	connection.Where("name = ?", user.Name).First(&dbuser)

	//check email is already registered or not
	if dbuser.Name != "" {
		var err Error
		err = SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	if user.Role == "user" {
		var err Error
		err = SetError(err, "Cannot create 'user' accounts")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return		
	}

	user.Password, err = GenerateHashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}

	//insert user details in database
	connection.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}