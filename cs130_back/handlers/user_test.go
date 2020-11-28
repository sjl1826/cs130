package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"database/sql/driver"
	"log"
	"time"
	"regexp"

	"github.com/gavv/httpexpect"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var columns = []string{"id", "first_name", "last_name", "email"}

func TestGetAllUsersSuccess(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)	
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, first_name, last_name, email FROM`)).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "Hunter", "Gunther", "henry@gmail.com"))

	obj :=	e.GET("/getAllUsers").
			Expect().
			Status(http.StatusOK).JSON().Object()
	

	//Testing the output
	var u UserSearchDetails
	u.ID = 1
	u.FirstName = "Hunter"
	u.LastName = "Gunther"
	u.Email = "henry@gmail.com"

	obj.Keys().ContainsOnly("users")
	obj.Value("users").Equal([]UserSearchDetails{u})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

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

	mux.HandleFunc("/getBuddiesListings", func(w http.ResponseWriter, r *http.Request) {
		GetBuddiesAndListings(db, w, r)
	})

	mux.HandleFunc("/createGroup", func(w http.ResponseWriter, r *http.Request) {
		CreateGroup(db, w, r)
	})

	mux.HandleFunc("/createListing", func(w http.ResponseWriter, r *http.Request) {
		CreateListing(db, w, r)
	})

	mux.HandleFunc("/sendInvitation", func(w http.ResponseWriter, r *http.Request) {
		SendInvitation(db, w, r)
	})

	mux.HandleFunc("/getAllUsers", func(w http.ResponseWriter, r *http.Request) {
		GetAllUsers(db, w, r)
	})

	return mux
}