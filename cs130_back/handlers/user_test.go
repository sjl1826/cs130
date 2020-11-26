package handlers

import (
	"net/http"
	"database/sql/driver"
	"log"
	"time"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)


//For Test Scenario (Create and Update Listing)
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

//	Setup functions

func GetMock() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New() //mock sql db
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	DB, err := gorm.Open("postgres", db) //open gorm db
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return mock, DB
}

func GetHandler(db *gorm.DB) http.Handler {
	
	mux := http.NewServeMux()

	mux.HandleFunc("/createListing", func(w http.ResponseWriter, r *http.Request) {
		CreateListing(db, w, r)
	})

	return mux
}