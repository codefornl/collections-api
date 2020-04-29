package models

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import the postgres dialect
)

// GormBase - Override the gorm.Model so we can hide fields we do not want
type GormBase struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// Link - HAL structure http://stateless.co/hal_specification.html#Links
type Link struct {
	Href string `json:"href,omitempty"`
}

// Filter - HAL structure
type Filter struct {
	Href      string `json:"href,omitempty"`
	Templated bool   `json:"templated,omitempty"`
}

// Cury - HAL structure http://stateless.co/hal_specification.html#CURIEs
type Cury struct {
	Href      string `json:"href,omitempty"`
	Name      string `json:"name,omitempty"`
	Templates bool   `json:"templates,omitempty"`
}

// ServiceLinks - Implementation of HAL for the index page
type ServiceLinks struct {
	Curies      []Cury `json:"curies"`
	Collections Link   `json:"cbases"`
	Usecases    Link   `json:"usecases"`
	Self        Link   `json:"cfnl:self"`
}

// Service - Primary Index response for HAL
type Service struct {
	Service     string       `json:"service"`
	About       string       `json:"about"`
	Browser     string       `json:"browser"`
	Application string       `json:"application"`
	Codebase    string       `json:"codebase"`
	Links       ServiceLinks `json:"_links"`
}

var db *gorm.DB //database

func init() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Category{}, &Collection{}, &Usecase{}, &Type{}, &Initiative{}) //Database migration
}

// GetDB - returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
