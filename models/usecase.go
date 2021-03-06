package models

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
	GormBase
	Collections  []Collection `gorm:"many2many:usecase_collection;"`
	Name         string       `json:"name"`
	Slug         string       `json:"slug"`
	Teaser       string       `json:"admin_name"`
	Description  string       `json:"description" gorm:"size:65531;not null"`
	Image        string       `json:"image" gorm:"not null"`
	TypeID       int
	Type         *Type
	LocationID   int
	Location     *Location
	Categories   []Category `json:"categories" gorm:"many2many:usecase_category;"`
	Organisation string     `json:"organisation" gorm:"not null"`
	Website      string     `json:"website" gorm:"not null"`
	Download     string     `json:"download" gorm:"not null"`
	License      string     `json:"license" gorm:"not null"`

	Links UsecaseLinks `json:"_links" gorm:"-"`
}

// Usecases - Plural HAL Response
type Usecases struct {
	Links    UsecasesLinks    `json:"_links"`
	Embedded UsecasesEmbedded `json:"_embedded"`
}
