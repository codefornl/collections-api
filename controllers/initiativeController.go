package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/codefornl/collections-api/models"
)

// GetInitiative - Get a single initiative
var GetInitiative = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)
	//vars := mux.Vars(r)
	//initiative := models.GetInitiative(vars["key"])
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Initiative{})
}

// GetInitiatives - Get multiple initiatives
var GetInitiatives = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)
	initiatives := models.GetInitiatives()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(initiatives)
}
