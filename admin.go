package main

import (
	"encoding/json"
	"net/http"
)

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)
	var newGame = Game{Active: true}
	connection.Model(&Game{}).Updates(Game{Active: false})
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

func NextTurn(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)
	var curGame Game
	var scores []Score
	connection.Model(&curGame).Where("active = ?", true).Association("Scores").DB.Order("updated_at desc").Find(&scores)
	var setNext bool
	for _, v := range scores {
		if v.Active {
			setNext = true
			connection.Model(&v).Updates(&Score{Active: false, Complete: true})
		}
		if setNext {
			connection.Model(&v).Updates(&Score{Active: true})
		}
	}
}

func HighScore(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var curGame Game
	var score Score
	var user User
	connection.Model(&curGame).Where("active = ?", true).Association("Scores").DB.Where("active = ?", true).First(&score)
	connection.Where("id = ?", true).First(&user)
	//gpio for highscore
	Initials(user.Initials)
	NextTurn(w, r)
}
