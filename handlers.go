package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	//get token from the route path
	vars := mux.Vars(r)
	token := vars["wsToken"]

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	connection := GetDatabase()
	defer CloseDatabase(connection)
	//get user by token
	var dbuser User
	connection.Where("wsToken = ?", token).First(&dbuser)
	//wsToken doesn't exist in db
	if dbuser.Name == "" {
		log.Print("wsToken does not exist in the database. No user was returned.")
		return
	}
	//single use token so clear it out
	dbuser.WsToken = ""
	connection.Save(&dbuser)

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		sw := string(message)
		switch sw {
		case "LU":
			Left(false)
		case "RU":
			Right(false)
		case "LD":
			Left(true)
		case "RD":
			Right(true)
		case "L":
			Launch()
		case "S":
			Start()
		}

		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	connection := GetDatabase()
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

	//check email is alredy registered or not
	if dbuser.Name != "" {
		var err Error
		err = SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}

	//insert user details in database
	connection.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	connection := GetDatabase()
	defer CloseDatabase(connection)

	var authDetails Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser User
	connection.Where("name = 	?", authDetails.Name).First(&authUser)

	if authUser.Name == "" {
		var err Error
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err Error
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := GenerateJWT(authUser.Name, authUser.Role)
	if err != nil {
		var err Error
		err = SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token Token
	token.Name = authUser.Name
	token.Role = authUser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, User."))
}

func GenerateUserWsUrl(w http.ResponseWriter, r *http.Request){
	connection := GetDatabase()
	defer CloseDatabase(connection)

	userName := r.Header.Get("Name")
	var dbuser User
	connection.Where("name = ?", userName).First(&dbuser)

	//username doesn't exist in db
	if dbuser.Name == "" {
		var err Error
		err = SetError(err, "Username not found in database")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	wsToken := uuid.New().String()
	dbuser.WsToken = wsToken
	connection.Save(&dbuser)

	//wsUrl += "wss:" hardcoded to http at the moment
	//create and return the websocket url with a user token in the route
	wsUrl := fmt.Sprintf("ws://%v/buttonpress/%v", r.Host, wsToken)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wsUrl)
}

