package main

import (
	"fmt"
	"github.com/codefornl/collections-api/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Setting up server, enabling CORS . . .")

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetServiceIndex).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/collection/{key}", controllers.GetCollection).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/cbases/{key}", controllers.GetCollection).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/collections", controllers.GetCollections).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/cbases", controllers.GetCollections).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/usecase/{key}", controllers.GetUsecase).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/usecases", controllers.GetUsecases).Methods(http.MethodGet, http.MethodOptions)
	router.Use(mux.CORSMethodMiddleware(router))

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}


	fmt.Println("Server is ready and is listening at port :" + port +" . . .")
	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
