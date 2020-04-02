package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Link struct {
	Href string `json:"href,omitempty"`
}

type Filter struct {
	Href string `json:"href,omitempty"`
	Templated bool `json:"templated,omitempty"`
}

type Cury struct {
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
	Templates bool `json:"templates,omitempty"`
}

type ServiceLinks struct {
	Curies []Cury `json:"curies"`
	Collections Link `json:"cbases"`
	Usecases Link `json:"usecases"`
	Self Link `json:"cfnl:self"`
}

type Service struct {
	Service string `json:"service"`
	About string `json:"about"`
	Browser string `json:"browser";gorm:"size:255"`
	Application string `json:"application";gorm:"size:255"`
	Codebase string `json:"codebase";gorm:"size:255"`
	Links ServiceLinks `json:"_links"`
}

var db *gorm.DB //database

func init() {

	conn, err := gorm.Open("sqlite3", "../collections-api.sqlite")
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Collection{},&Usecase{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}