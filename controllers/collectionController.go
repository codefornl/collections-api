package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codefornl/collections-api/models"
	"github.com/gorilla/mux"
)

// GetCollectionsLinks - Get the links that should be added under `_links` in the HAL response
var GetCollectionsLinks = func(r *http.Request) models.CollectionsLinks {
	return models.CollectionsLinks{
		Self: models.Link{
			Href: GetSelf(r),
		},
		Home: models.Link{
			Href: GetBase(r),
		},
		Filter: GetFilter(r),
	}
}

// GetCollection - Get a single collection
var GetCollection = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)
	vars := mux.Vars(r)
	collection := &models.Collection{}
	models.GetDB().Find(&collection, vars["key"])
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Collection{})
}

// GetCollections - Get multiple collections
var GetCollections = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)

	collections := []models.Collection{}
	models.GetDB().Find(&collections)
	response := models.Collections{
		Links: GetCollectionsLinks(r),
		Embedded: models.CollectionsEmbedded{
			Collections: collections,
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateCollection - Create
var CreateCollection = func(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	var e models.Collection
	err := decoder.Decode(&e)
	if err != nil {
		fmt.Print(err)
		return
	}
	models.GetDB().Create(&e)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{"message": "Success"}`)
}
