package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func SocketButton(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		sw := string(message)
		if params["userID"] != activeUser {
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
		}
		select {
		case usr := <-currentUser:
			if usr != activeUser {
				activeUser = usr
				c.WriteMessage(mt, []byte("NEW TURN"))
			}
		default:
			fmt.Println("no message received")
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		JSONError(w, err, 500)
		return
	}

	var dbuser User
	var curGame Game
	// var scores []Score
	connection.Where("name = ?", user.Name).First(&dbuser)

	//check email is alredy registered or not
	if dbuser.Name != "" {
		var err Error
		err = SetError(err, "Name already in use")
		JSONError(w, err, 500)
		return
	}

	user.Password, err = GenerateHashPassword(user.Password)
	if err != nil {
		JSONError(w, err, 500)
	}

	// insert user details in database
	connection.Model(&curGame).Where("in_active = ?", false).First(&curGame)
	connection.Model(&curGame).Where("in_active = ?", false).Association("Users").Append(&user)
	w.Header().Set("Content-Type", "application/json")
	// get list of users to return as queued players
	var users []User
	var scores []Score
	connection.Model(&curGame).Order("updated_at desc").Association("Users").Find(&users)
	connection.Model(&curGame).Order("updated_at desc").Association("Scores").Find(&scores)
	if len(scores) < 1 {
		connection.Model(&curGame).Where("in_active = ?", false).Association("Scores").Append(&Score{User: user.ID, Active: true})
		currentUser <- user.ID.String()
	} else {
		connection.Model(&curGame).Where("in_active = ?", false).Association("Scores").Append(&Score{User: user.ID})
	}
	var players []Player

	for _, element := range users {
		players = append(players, Player{Name: element.Name, Initials: element.Initials, Class: GetScoreClass(scores, element.ID), Score: GetScoreValue(scores, element.ID)})
	}
	// json.NewEncoder(w).Encode(players)
	validToken, err := GenerateUserJWT(user.Name, user.ID.String())
	if err != nil {
		var err Error
		err = SetError(err, "Failed to generate token")
		JSONError(w, err, 500)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	http.SetCookie(w, &http.Cookie{
		Name:    "Authorization",
		Value:   "Bearer " + validToken,
		Expires: expirationTime,
	})
	var token Token
	token.Name = user.Name
	token.ID = user.ID.String()
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func GetCurrentGame(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var curGame Game
	var scores []Score

	connection.Model(&curGame).Where("in_active = ?", false).First(&curGame)
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)

	// get list of users to return as queued players
	var users []User
	connection.Model(&curGame).Order("created_at asc").Association("Users").Find(&users)
	connection.Model(&curGame).Order("created_at asc").Association("Scores").Find(&scores)
	var players []Player
	for _, element := range users {
		players = append(players, Player{Name: element.Name, Initials: element.Initials, Class: GetScoreClass(scores, element.ID), Score: GetScoreValue(scores, element.ID)})
	}

	json.NewEncoder(w).Encode(players)
}

func GetCurrentGameWID(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var curGame Game
	var scores []Score

	connection.Model(&curGame).Where("in_active = ?", false).First(&curGame)
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)

	// get list of users to return as queued players
	var users []User
	connection.Model(&curGame).Order("created_at asc").Association("Users").Find(&users)
	connection.Model(&curGame).Order("created_at asc").Association("Scores").Find(&scores)
	var players []AdminPlayer
	for _, element := range users {
		players = append(players, AdminPlayer{ID: element.ID.String(), Name: element.Name, Initials: element.Initials, Class: GetScoreClass(scores, element.ID), Score: GetScoreValue(scores, element.ID)})
	}

	json.NewEncoder(w).Encode(players)
}

func HandleQueue() {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)
}

//signin in the same action for users and admins but with different secretkey credentials
func PostSignIn(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var authDetails Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		JSONError(w, err, 500)
		return
	}

	var authUser User
	connection.Where("name = 	?", authDetails.Name).First(&authUser)

	if authUser.Name == "" {
		var err Error
		err = SetError(err, "Username or Password is incorrect")
		JSONError(w, err, 403)
		return
	}

	check := CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err Error
		err = SetError(err, "Username or Password is incorrect")
		JSONError(w, err, 403)
		return
	}

	validToken, err := GenerateUserJWT(authUser.Name, authUser.ID.String())
	if err != nil {
		var err Error
		err = SetError(err, "Failed to generate token")
		JSONError(w, err, 500)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	http.SetCookie(w, &http.Cookie{
		Name:    "Authorization",
		Value:   "Bearer " + validToken,
		Expires: expirationTime,
	})
	var token Token
	token.Name = authUser.Name
	token.ID = authUser.ID.String()
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func GetListControls(w http.ResponseWriter, r *http.Request) {
	var controls []Control
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	result := connection.Find(&controls)
	if result.Error != nil {
		var err Error
		err = SetError(err, "Failed to get controls from the db")
		JSONError(w, err, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controls)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	//vars := mux.Vars(r)
	//userId := vars["userId"]

	userId := r.URL.Query().Get("userId")
	role := r.URL.Query().Get("role")

	if role == "admin" {
		var dbAdmin Admin
		connection.Where("id = ?", userId).First(&dbAdmin)

		if dbAdmin.Email == "" {
			var err Error
			err = SetError(err, "Username does't exist")
			JSONError(w, err, 500)
			return
		}

		//can't delete self
		if r.Header.Get("Email") == dbAdmin.Email {
			var err Error
			err = SetError(err, "User can't delete themselves")
			JSONError(w, err, 500)
			return
		}
		connection.Delete(&dbAdmin)
	} else {
		var dbUser User
		connection.Where("id = ?", userId).First(&dbUser)

		if dbUser.Name == "" {
			var err Error
			err = SetError(err, "Username does't exist")
			JSONError(w, err, 500)
			return
		}

		//can't delete self
		if r.Header.Get("Name") == dbUser.Name {
			var err Error
			err = SetError(err, "User can't delete themselves")
			JSONError(w, err, 500)
			return
		}
		connection.Delete(&dbUser)
	}
}

func PostLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "Authorization",
		MaxAge: -1,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Status: 200, Message: "logged out."})
}
