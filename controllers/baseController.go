package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/codefornl/collections-api/models"
)

// PreflightCheck - Check for CORS request. This one is essential for Angular to function
func PreflightCheck(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	if r.Method == http.MethodOptions {
		return
	}
}

// GetBase - Get the information on the service host
var GetBase = func(r *http.Request) string {
	scheme := "http"
	path := ""
	if r.TLS != nil {
		scheme = "https"
	}
	if r.Header.Get("X-Original-URI") != "" {
		path = r.Header.Get("X-Original-URI")
		// strip trailing slash
		path = strings.TrimSuffix(path, "/")
	}

	return scheme + "://" + r.Host + path
}

// GetSelf - Get information about the request url used
var GetSelf = func(r *http.Request) string {
	return GetBase(r) + r.URL.Path
}

// GetFilter - Construct a filter object from the request to facilitate the response
var GetFilter = func(r *http.Request) models.Filter {
	return models.Filter{
		Templated: true,
		Href:      GetSelf(r) + "{?q}",
	}
}

// GetServiceIndex - Get the HAL content for the API Entry point. Should provide information about
// the underlying routes
var GetServiceIndex = func(w http.ResponseWriter, r *http.Request) {
	PreflightCheck(w, r)
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
				{
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
