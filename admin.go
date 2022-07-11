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

func NextTurn(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)
	var curGame Game
	var scores []Score
	connection.Model(&curGame).Where("in_active = ?", false).Order("updated_at desc").Association("Scores").Find(&scores)
	var setNext bool
	for _, v := range scores {
		if v.Active && !v.Complete {
			setNext = true
			connection.Model(&v).Updates(&Score{Complete: true})
		}
		if setNext {
			connection.Model(&v).Updates(&Score{Active: true})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scores)
}

func UpdateScore(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var newScore ScoreUpdate

	err := json.NewDecoder(r.Body).Decode(&newScore)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	var curGame Game
	var scores []Score
	connection.Model(&curGame).Where("in_active = ?", false).First(&curGame)
	connection.Model(&curGame).Where("in_active = ?", false).Order("updated_at desc").Association("Scores").Find(&scores)
	for _, v := range scores {
		if v.User == newScore.ID {
			connection.Model(&v).Updates(&Score{Score: int64(newScore.Score)})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scores)
}

func HighScore(w http.ResponseWriter, r *http.Request) {
	connection, _ := GetDatabase()
	defer CloseDatabase(connection)

	var curGame Game
	var scores []Score
	var user User
	connection.Model(&curGame).Where("in_active = ?", false).Order("updated_at desc").Association("Scores").Find(&scores)
	for _, v := range scores {
		if v.Active && !v.Complete {
			connection.Where("id = ?", v.User).First(&user)
			connection.Model(&v).Updates(&Score{Complete: true})
		}
	}
	//gpio for highscore
	Initials(user.Initials)
	NextTurn(w, r)
}
