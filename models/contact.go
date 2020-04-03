package models

// Contact - A person that `owns` the object. Not persé a user in the system
type Contact struct {
	GormBase
	Name        string `json:"name" gorm:"not null"`
	Avatar      string `json:"avatar,omitempty"`
	Role        string `json:"role,omitempty"`
	Phonenumber string `json:"phonenumber,omitempty"`
	Email       string `json:"email,omitempty"`
}
