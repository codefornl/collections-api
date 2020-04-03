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

// GetInitiative - Singular DB Getter
func GetInitiative(u string) Initiative {
	usecase := Initiative{}
	GetDB().Find(&usecase, 1)
	return usecase
}

// GetInitiatives - Plural DB Getter
func GetInitiatives() []Initiative {
	usecases := []Initiative{}
	GetDB().Find(&usecases)
	return usecases
}
