package main

import (
	"github.com/gorilla/mux"
)

//--------GLOBAL VARIABLES---------------

var (
	router    *mux.Router
	secretkey string = "secretkeyjwt"
)

func main() {
	// go Blink()
	InitialMigration()
	CreateRouter()
	InitializeRoute()
	InitializeStatic()
	ServerStart()
}
