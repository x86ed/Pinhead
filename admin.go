package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}

func PostNewGame(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)
	var newGame = Game{}
	connection.Model(&Game{}).Where("1 == 1").Updates(Game{InActive: true})
	connection.Create(&newGame)
	if newGame.ID.String() == "" {
		var err Error
		err = SetError(err, "Couldn't create a new Game")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newGame)
}

func PostNextTurn(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)
	var curGame Game
	var scores []Score
	connection.Model(&curGame).Where("in_active = ?", false).First(&curGame)
	connection.Model(&curGame).Order("updated_at desc").Association("Scores").Find(&scores)
	for _, v := range scores {
		if v.Active && !v.Complete {
			connection.Model(&v).Updates(map[string]interface{}{"complete": true, "active": false})
		}
		if !v.Active && !v.Complete {
			currentUser <- v.User.String()
			connection.Model(&v).Updates(Score{Active: true})
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("scores: %+v\n", scores)
	json.NewEncoder(w).Encode(scores)
}

func PostUpdateScore(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var newScore ScoreUpdate

	err := json.NewDecoder(r.Body).Decode(&newScore)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		JSONError(w, err, 500)
		return
	}
	var curGame Game
	var scores []Score
	connection.Model(&curGame).Where("in_active = ?", false).First(&curGame)
	connection.Model(&curGame).Order("updated_at desc").Association("Scores").Find(&scores)
	for _, v := range scores {
		if v.User.String() == newScore.ID {
			connection.Model(&v).Updates(&Score{Score: int64(newScore.Score)})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("scores: %+v\n", scores)
	json.NewEncoder(w).Encode(scores)
}

func PostHighScore(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var curGame Game
	var scores []Score
	var user User
	connection.Model(&curGame).Where("active = ?", false).Order("updated_at desc").Association("Scores").Find(&scores)
	for _, v := range scores {
		if v.Active && !v.Complete {
			connection.Where("id = ?", v.User).First(&user)
			connection.Model(&v).Updates(&Score{Complete: true})
		}
	}
	//gpio for highscore
	Initials(user.Initials)
	PostNextTurn(w, r)
}

func PostCreateAdmin(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var user Admin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		JSONError(w, err, 500)
		return
	}

	var dbuser Admin
	connection.Where("name = ?", user.Email).First(&dbuser)

	//check email is already registered or not
	if dbuser.Email != "" {
		var err Error
		err = SetError(err, "Email already in use")
		JSONError(w, err, 500)
		return
	}

	user.Password, err = GenerateHashPassword(user.Password)
	if err != nil {
		JSONError(w, err, 500)
		return
	}

	//insert user details in database
	connection.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetListUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	result := connection.Find(&users)

	if result.Error != nil {
		var err Error
		err = SetError(err, "Failed to get users from the db")
		JSONError(w, err, 500)
		return
	}

	for _, user := range users {
		user.MarshalJSON()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (u User) MarshalJSON() ([]byte, error) {
	// prevent recursion
	type user User
	x := user(u)
	// remove users password so it is not returned to the caller
	x.Password = ""
	return json.Marshal(x)
}

func GetListAdmins(w http.ResponseWriter, r *http.Request) {
	var admins []Admin

	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	result := connection.Find(&admins)

	if result.Error != nil {
		var err Error
		err = SetError(err, "Failed to get users from the db")
		JSONError(w, err, 500)
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

func PostAdminSignIn(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var authDetails AdminAuthentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		JSONError(w, err, 500)
		return
	}

	var authUser Admin
	connection.Where("email = 	?", authDetails.Email).First(&authUser)

	if authUser.Email == "" {
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

	validToken, err := GenerateAdminJWT(authUser.Email)
	if err != nil {
		var err Error
		err = SetError(err, "Failed to generate token")
		JSONError(w, err, 500)
		return
	}

	var token AdminToken
	token.Email = authUser.Email
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
