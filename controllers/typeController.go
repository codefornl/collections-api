package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codefornl/collections-api/models"
)

// CreateType - Create
var CreateType = func(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	var e models.Type
	err := decoder.Decode(&e)
	if err != nil {
		fmt.Print(err)
		return
	}
	models.GetDB().Create(&e)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{"message": "Success"}`)
}
