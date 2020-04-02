package main

import (
	"fmt"
	"github.com/codefornl/collections-api/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetServiceIndex).Methods("GET")
	router.HandleFunc("/collection/{key}", controllers.GetCollection).Methods("GET")
	router.HandleFunc("/collections", controllers.GetCollections).Methods("GET")
	router.HandleFunc("/usecase/{key}", controllers.GetUsecase).Methods("GET")
	router.HandleFunc("/usecases", controllers.GetUsecases).Methods("GET")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
