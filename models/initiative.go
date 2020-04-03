package models

import (
	"github.com/jinzhu/gorm"
)

// Initiative model taken from the helpradar datamodel:
//
// ParentInitiativeId -> Link to other Initiative
// Name
// Type -> Link to Types
type Initiative struct {
	gorm.Model
	Name string `json:"name"`
}
