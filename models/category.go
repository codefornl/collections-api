package models

// Category of initiative. Can be used to group initiatives together,
// and better define related initiatives
// Examples taken from the helpradar datamodel:
// - Transport
// - Healthcare
// - Care
// - Medical Supplies
// - Food
// - Non Food
// - Social
// - Psychological
// - Entertainment
type Category struct {
	GormBase
	Name string `json:"name"`
}
