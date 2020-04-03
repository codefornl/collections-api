package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/codefornl/collections-api/models"
	"github.com/gorilla/mux"
)

// GetInitiative - Get a single initiative
var GetInitiative = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)
	vars := mux.Vars(r)
	initiative := &models.Initiative{}
	models.GetDB().Find(&initiative, vars["key"])
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(initiative)
}

// GetInitiatives - Get multiple initiatives
var GetInitiatives = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)
	initiatives := []models.Initiative{}
	models.GetDB().Find(&initiatives)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(initiatives)
}
