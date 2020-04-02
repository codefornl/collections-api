package models

import (
	"github.com/jinzhu/gorm"
)

// Singular HAL Links
type CollectionLinks struct {
	Self Link `json:"self,omitempty"`
	Home Link `json:"home,omitempty"`
	Slug Link `json:"self_slug,omitempty"`
	Filter Filter `json:"filter,omitempty"`
	// cbases
}

// Plural HAL Links
type CollectionsLinks struct {
	Self Link `json:"self,omitempty"`
	Home Link `json:"home,omitempty"`
	Filter Filter `json:"filter,omitempty"`
}

// Plural HAL Embedded
type CollectionsEmbedded struct {
	Collections []Collection `json:"cbase,omitempty"`
}

// Basic ORM, has Singular HAL Response
type Collection struct {
	gorm.Model `json:"-"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
	AdminName string `json:"admin_name,omitempty";gorm:"size:255"`
	AdminEmail string `json:"admin_email,omitempty";gorm:"size:255;not null"`
	TokenEncrypted string `json:"token_encrypted,omitempty";gorm:"size:255;not null"`
	Image string `json:"image,omitempty";gorm:"size:255;not null"`
	Description string `json:"description,omitempty";gorm:"size:65531;not null"`
	Language string `json:"language,omitempty";gorm:"size:3";not null"`
	Promote bool `json:"-";gorm:"not null"`
	LogoImage string `json:"logo_image,omitempty";gorm:"size:255;not null"`
	HighlightColor string `json:"highlight_color,omitempty";gorm:"size:7;not null"`
	Disables bool `json:"-";gorm:"not null"`
	Links CollectionLinks `json:"_links,omitempty";gorm:"-"`
}

// Plural HAL Response
type Collections struct {
	Links CollectionsLinks `json:"_links,omitempty"`
	Embedded CollectionsEmbedded `json:"_embedded,omitempty"`
}

// Singular DB Getter
func GetCollection(u string) *Collection {
	collection := &Collection{}
	GetDB().Find(&collection,1)
	return collection
}

// Plural DB Getter
func GetCollections() []Collection {
	collections := []Collection{}
	GetDB().Find(&collections)
	return collections
}

