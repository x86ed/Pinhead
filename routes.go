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
	router.HandleFunc("/signup", PostSignUp).Methods("POST")
	router.HandleFunc("/signin", PostSignIn).Methods("POST")

	router.HandleFunc("/controls", GetListControls).Methods("GET")
	router.HandleFunc("/game", IsAuthorizedUser(GetCurrentGame)).Methods("GET")

	// Websocket route
	router.HandleFunc("/buttonpress/{userID}", SocketButton).Methods("GET")
	router.HandleFunc("/logout", IsAuthorizedUser(PostLogout)).Methods("POST")
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})

	//local server for admin
    localRouter.HandleFunc("/players", IsAuthorizedAdmin(ListActivePlayers)).Methods("GET")
	localRouter.HandleFunc("/users", IsAuthorizedAdmin(GetListUsers)).Methods("GET")
	localRouter.HandleFunc("/admins", IsAuthorizedAdmin(GetListAdmins)).Methods("GET")

	localRouter.HandleFunc("/new_game", IsAuthorizedAdmin(PostNewGame)).Methods("POST")
	localRouter.HandleFunc("/next_turn", IsAuthorizedAdmin(PostNextTurn)).Methods("POST")
	localRouter.HandleFunc("/high_score", IsAuthorizedAdmin(PostHighScore)).Methods("POST")
	localRouter.HandleFunc("/update_score", IsAuthorizedAdmin(PostUpdateScore)).Methods("POST")

	localRouter.HandleFunc("/signin", PostAdminSignIn).Methods("POST")
	localRouter.HandleFunc("/admin", IsAuthorizedAdmin(PostCreateAdmin)).Methods("POST")
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
		fmt.Println("Local Admin Server started at http://localhost:54321")
		http.ListenAndServe(":54321", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type"}), handlers.AllowedMethods([]string{"POST", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:54321"}))(localRouter))
	}()

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		return err
	}
	return nil
}
