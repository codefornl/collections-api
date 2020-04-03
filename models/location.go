package models

// Location - what more can I say?
type Location struct {
	GormBase
	Name          string  `json:"name"`
	Address       string  `json:"address,omitempty"`
	Neighbourhood string  `json:"neighbourhood,omitempty"`
	Municipality  string  `json:"municipality,omitempty"`
	Region        string  `json:"region,omitempty"`
	Country       string  `json:"country,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
}
