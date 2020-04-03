package models

import (
	"github.com/jinzhu/gorm"
)

// UsecaseLinks - Singular HAL Links
type UsecaseLinks struct {
	Self Link `json:"self"`
	Home Link `json:"home"`
	Slug Link `json:"self_slug"`
	// usecases
}

// UsecasesLinks - Plural HAL Links
type UsecasesLinks struct {
	Self Link `json:"self"`
	Home Link `json:"home"`
}

// UsecasesEmbedded - Plural HAL Embedded
type UsecasesEmbedded struct {
	Usecases []Usecase `json:"usecase"`
}

// Usecase - Basic ORM, has Singular HAL Response
type Usecase struct {
	gorm.Model   `json:"-"`
	Collections  []Collection `gorm:"many2many:usecase_collection;"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Slug         string       `json:"slug"`
	Teaser       string       `json:"admin_name"`
	Description  string       `json:"description" gorm:"size:65531;not null"`
	Image        string       `json:"image" gorm:"not null"`
	Type         string       `json:"type" gorm:"not null"`
	Location     string       `json:"location" gorm:"not null"`
	Country      string       `json:"country" gorm:"not null"`
	Category     string       `json:"category" gorm:"not null"`
	Organisation string       `json:"organisation" gorm:"not null"`
	Website      string       `json:"website" gorm:"not null"`
	Download     string       `json:"download" gorm:"not null"`
	License      string       `json:"license" gorm:"not null"`
	ContactName  string       `json:"contact_name" gorm:"not null"`
	ContactImage string       `json:"contact_image" gorm:"not null"`
	ContactEmail string       `json:"contact_email" gorm:"not null"`
	Links        UsecaseLinks `json:"_links" gorm:"-"`
}

// Usecases - Plural HAL Response
type Usecases struct {
	Links    UsecasesLinks    `json:"_links"`
	Embedded UsecasesEmbedded `json:"_embedded"`
}

// GetUsecase - Singular DB Getter
func GetUsecase(u string) Usecase {
	usecase := Usecase{}
	GetDB().Find(&usecase, 1)
	return usecase
}

// GetUsecases - Plural DB Getter
func GetUsecases() []Usecase {
	usecases := []Usecase{}
	GetDB().Find(&usecases)
	return usecases
}
