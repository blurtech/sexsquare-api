package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"sexsquare-api/app"
	"sexsquare-api/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	err := http.ListenAndServe(":"+port, handlers.CORS()(loggedRouter)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
