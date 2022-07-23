package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
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
	var curGame Game
	var scores []Score
	connection.Where("name = ?", user.Name).First(&dbuser)

	//check email is alredy registered or not
	if dbuser.Name != "" {
		var err Error
		err = SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	if user.Role == "admin" {
		var err Error
		err = SetError(err, "Cannot create'admin' accounts")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = GenerateHashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}

	// insert user details in database
	connection.Create(&user)
	connection.Model(&curGame).Where("in_active = ?", false).First(&curGame)
	connection.Where("name = ?", user.Name).First(&dbuser)
	connection.Model(&curGame).Where("in_active = ?", false).Association("Users").Append(&user)
	connection.Model(&curGame).Where("in_active = ?", false).Association("Scores").Append(&Score{User: user.ID})
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)

	// get list of users to return as queued players
	var users []User
	connection.Model(&curGame).Order("updated_at desc").Association("Users").Find(&users)
	connection.Model(&curGame).Order("updated_at desc").Association("Scores").Find(&scores)
	var players []Player

	for _, element := range users {
		players = append(players, Player{Name: element.Name, Initials: element.Initials, Class: GetScoreState(scores, element.ID)})
	}

	json.NewEncoder(w).Encode(players)
}

func HandleQueue() {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)
}

func AdminSignIn(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var authDetails AdminAuthentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser Admin
	connection.Where("email = 	?", authDetails.Email).First(&authUser)

	if authUser.Email == "" {
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

	validToken, err := GenerateAdminJWT(authUser.Email)
	if err != nil {
		var err Error
		err = SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token AdminToken
	token.Email = authUser.Email
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

//signin in the same action for users and admins but with different secretkey credentials
func UserSignIn(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
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

	validToken, err := GenerateUserJWT(authUser.Name, authUser.Role)
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

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, User."))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	result := connection.Find(&users)

	if result.Error != nil {
		var err Error
		err = SetError(err, "Failed to get users from the db")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	for _, user := range users {
		user.MarshalJSON()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func ListAdmins(w http.ResponseWriter, r *http.Request) {
	var admins []Admin

	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	result := connection.Find(&admins)

	if result.Error != nil {
		var err Error
		err = SetError(err, "Failed to get users from the db")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	for _, admin := range admins {
		admin.MarshalJSON()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admins)
}

func (a Admin) MarshalJSON() ([]byte, error) {
	// prevent recursion
	type admin Admin
	x := admin(a)
	// remove users password so it is not returned to the caller
	x.Password = ""
	return json.Marshal(x)
}

func (u User) MarshalJSON() ([]byte, error) {
	// prevent recursion
	type user User
	x := user(u)
	// remove users password so it is not returned to the caller
	x.Password = ""
	return json.Marshal(x)
}

func CreateAdminAccount(w http.ResponseWriter, r *http.Request) {
	CreateAdmin(w, r)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	//vars := mux.Vars(r)
	//userId := vars["userId"]

	userId := r.URL.Query().Get("userId")
	role := r.URL.Query().Get("role")

	if (role == "admin") {
		var dbAdmin Admin
		connection.Where("id = ?", userId).First(&dbAdmin)
	
		if dbAdmin.Email == "" {
			var err Error
			err = SetError(err, "Username does't exist")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
	
		//can't delete self
		if r.Header.Get("Email") == dbAdmin.Email {
			var err Error
			err = SetError(err, "User can't delete themselves")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
		connection.Delete(&dbAdmin)
	} else {
		var dbUser User
		connection.Where("id = ?", userId).First(&dbUser)
	
		if dbUser.Name == "" {
			var err Error
			err = SetError(err, "Username does't exist")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
	
		//can't delete self
		if r.Header.Get("Name") == dbUser.Name {
			var err Error
			err = SetError(err, "User can't delete themselves")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
		connection.Delete(&dbUser)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//JWT tokens typically just expire
	//are we going to implement something like cookies instead that we can revoke?
}
