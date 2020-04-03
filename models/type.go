package models

import (
	"github.com/jinzhu/gorm"
)

// Type of initiative. Can be used to group initiatives together.
// Examples taken from the helpradar datamodel:
// - Religous Centre
// - Sports Club
// - Community Center
// - Nursing Home
// - Commercial
// - NGO's (This are pro and possibly not belong in this list.)
// - Online Platform
type Type struct {
	gorm.Model
	Name string `json:"name"`
}
