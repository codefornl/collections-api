package models

import (
	"github.com/jinzhu/gorm"
)

// Singular HAL Links
type UsecaseLinks struct {
	Self Link `json:"self"`
	Home Link `json:"home"`
	Slug Link `json:"self_slug"`
	// usecases
}

// Plural HAL Links
type UsecasesLinks struct {
	Self Link `json:"self"`
	Home Link `json:"home"`
}

// Plural HAL Embedded
type UsecasesEmbedded struct {
	Usecases []Usecase `json:"usecase"`
}

// Basic ORM, has Singular HAL Response
type Usecase struct {
	gorm.Model `json:"-"`
	Collections []Collection `gorm:"many2many:usecase_collection;"`
	ID string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Teaser string `json:"admin_name";gorm:"size:255"`
	Description string `json:"description";gorm:"size:65531;not null"`
	Image string `json:"image";gorm:"size:255;not null"`
	Type string `json:"type";gorm:"size:255;not null"`
	Location string `json:"location";gorm:"size:255;not null"`
	Country string `json:"country";gorm:"size:255;not null"`
	Category string `json:"category";gorm:"size:255;not null"`
	Organisation string `json:"organisation";gorm:"size:255;not null"`
	Website string `json:"website";gorm:"size:255;not null"`
	Download string `json:"download";gorm:"size:255;not null"`
	License string `json:"license";gorm:"size:255;not null"`
	ContactName string `json:"contact_name";gorm:"size:255;not null"`
	ContactImage string `json:"contact_image";gorm:"size:255;not null"`
	ContactEmail string `json:"contact_email";gorm:"size:255;not null"`
	Links UsecaseLinks `json:"_links";gorm:"-"`
}

// Plural HAL Response
type Usecases struct {
	Links UsecasesLinks `json:"_links"`
	Embedded UsecasesEmbedded `json:"_embedded"`
}

// Singular DB Getter
func GetUsecase(u string) Usecase {
	usecase := Usecase{}
	GetDB().Find(&usecase, 1)
	return usecase
}

// Plural DB Getter
func GetUsecases() []Usecase {
	usecases := []Usecase{}
	GetDB().Find(&usecases)
	return usecases
}