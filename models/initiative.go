package models

// Initiative model taken from the helpradar datamodel:
type Initiative struct {
	GormBase
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	ParentID   int
	Parent     *Initiative
	TypeID     int
	Type       *Type
	Categories []Category `gorm:"many2many:initiative_category;"`
	Website    string     `json:"website" gorm:"not null"`
	LocationID int
	Location   *Location
	Contacts   []Contact `gorm:"many2many:initiative_contact;"`
}
