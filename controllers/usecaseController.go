package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codefornl/collections-api/models"
	"github.com/gorilla/mux"
)

// GetUsecasesLinks - Get the links that should be added under `_links` in the HAL response
var GetUsecasesLinks = func(r *http.Request) models.UsecasesLinks {
	return models.UsecasesLinks{
		Self: models.Link{
			Href: GetSelf(r),
		},
		Home: models.Link{
			Href: GetBase(r),
		},
	}
}

//GetUsecase - Get a single usecase
var GetUsecase = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)
	vars := mux.Vars(r)
	usecase := &models.Usecase{}
	models.GetDB().Find(&usecase, vars["key"])
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usecase)
}

// GetUsecases - Get multiple usecases
var GetUsecases = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)
	usecases := []models.Usecase{}
	models.GetDB().Find(&usecases)
	response := models.Usecases{
		Links: GetUsecasesLinks(r),
		Embedded: models.UsecasesEmbedded{
			Usecases: usecases,
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateUsecase - Create
var CreateUsecase = func(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	var e models.Usecase
	err := decoder.Decode(&e)
	if err != nil {
		fmt.Print(err)
		return
	}
	models.GetDB().Create(&e)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{"message": "Success"}`)
}
