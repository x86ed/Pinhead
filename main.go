//go:generate ./scripts/make-env.sh
package main

import (
	"log"

	"github.com/gorilla/mux"
)

//--------GLOBAL VARIABLES---------------

var (
	router            *mux.Router
	localRouter       *mux.Router
	userSecretKey, _  = GoDotEnvVariable("JWTKEY")
	adminSecretKey, _ = GoDotEnvVariable("ADMINJWTKEY")
	adminUser, _      = GoDotEnvVariable("ADMINEMAIL")
	adminPass, _      = GoDotEnvVariable("ADMINPASS")
	activeUser        string
)

func main() {
	err := InitialMigration()
	if err != nil {
		log.Fatal("Couldn't initalize service")
	}
	CreateRouter()
	InitializeRoute()
	InitializeStatic()
	err = ServerStart()
	if err != nil {
		log.Fatal("Couldn't initalize service")
	}
}
