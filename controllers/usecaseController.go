package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/codefornl/collections-api/models"
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
	//vars := mux.Vars(r)
	//usecase := models.GetUsecase(vars["key"])
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Usecase{})
}

// GetUsecases - Get multiple usecases
var GetUsecases = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)
	usecases := models.GetUsecases()
	response := models.Usecases{
		Links: GetUsecasesLinks(r),
		Embedded: models.UsecasesEmbedded{
			Usecases: usecases,
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
