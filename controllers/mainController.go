package controllers

import (
	"encoding/json"
	"github.com/codefornl/collections-api/models"
	"net/http"
	//"github.com/gorilla/mux"
)

var GetBase = func(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

var GetSelf = func(r *http.Request) string {
	return GetBase(r) + r.URL.Path
}

var GetFilter = func(r *http.Request) models.Filter {
	return models.Filter{
		Templated: true,
		Href:      GetSelf(r) + "{?q}",
	}
}

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

var GetServiceIndex = func(w http.ResponseWriter, r *http.Request) {
	application := GetSelf(r)
	about := "https://www.codefor.nl/clarity"
	service := models.Service{
		Service:     "cbase: curated sets of case studies of digital tools for government",
		Browser:     "https://haltalk.herokuapp.com/explorer/browser.html#" + application,
		Application: application,
		Codebase:    "https://github.com/codefornl/collections-api",
		About:       about,
		Links: models.ServiceLinks{
			Curies: []models.Cury{
				models.Cury{
					Name:      "cfnl",
					Templates: false,
					Href:      about,
				},
			},
			Usecases: models.Link{
				Href: application + "usecases",
			},
			Collections: models.Link{
				Href: application + "collections",
			},
			Self: models.Link{
				Href: application,
			},
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

var GetCollection = func(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//collection := models.GetCollection(vars["key"])
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Collection{})
}

var GetCollections = func(w http.ResponseWriter, r *http.Request) {
	collections := models.GetCollections()
	response := models.Collections{
		Links: GetCollectionsLinks(r),
		Embedded: models.CollectionsEmbedded{
			Collections: collections,
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

var GetUsecase = func(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//usecase := models.GetUsecase(vars["key"])
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Usecase{})
}

var GetUsecases = func(w http.ResponseWriter, r *http.Request) {
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
