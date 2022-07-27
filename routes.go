package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// create a mux router
func CreateRouter() {
	router = mux.NewRouter()
	localRouter = mux.NewRouter()
}

// static files
func InitializeStatic() {
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	localRouter.PathPrefix("/").Handler(http.FileServer((http.Dir("./AdminPage/"))))
}

// initialize all routes
func InitializeRoute() {
	router.HandleFunc("/signup", SignUp).Methods("POST")
	router.HandleFunc("/signin", UserSignIn).Methods("POST")

	router.HandleFunc("/user", IsAuthorizedUser(UserIndex)).Methods("GET")
	router.HandleFunc("/controls", ListControls).Methods("GET")

	router.HandleFunc("/buttonpress", echo).Methods("GET")
	router.HandleFunc("/logout", IsAuthorizedUser(Logout)).Methods("POST")
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})

	//local server for admin
	localRouter.HandleFunc("/users", IsAuthorizedAdmin(ListUsers)).Methods("GET")
	localRouter.HandleFunc("/players", IsAuthorizedAdmin(ListActivePlayers)).Methods("GET")
	localRouter.HandleFunc("/admins", IsAuthorizedAdmin(ListAdmins)).Methods("GET")

	localRouter.HandleFunc("/new_game", IsAuthorizedAdmin(NewGame)).Methods("POST")
	localRouter.HandleFunc("/next_turn", IsAuthorizedAdmin(NextTurn)).Methods("POST")
	localRouter.HandleFunc("/high_score", IsAuthorizedAdmin(HighScore)).Methods("POST")
	localRouter.HandleFunc("/update_score", IsAuthorizedAdmin(UpdateScore)).Methods("POST")

	localRouter.HandleFunc("/signin", AdminSignIn).Methods("POST")
	localRouter.HandleFunc("/admin", IsAuthorizedAdmin(CreateAdmin)).Methods("POST")
	localRouter.HandleFunc("/admin", IsAuthorizedAdmin(CreateAdminAccount)).Methods("POST")
	localRouter.HandleFunc("/admin", IsAuthorizedAdmin(DeleteAccount)).Methods("DELETE")

	localRouter.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}

// start the server
func ServerStart() error {

	go func() {
		//this server is only available on local host to allow for adding an admin user
		//don't change the ip address from localhost
		fmt.Println("Local Admin Server started at http://localhost:44321")
		http.ListenAndServe(":44321", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type"}), handlers.AllowedMethods([]string{"POST", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:44321"}))(localRouter))
	}()

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		return err
	}
	return nil
}
