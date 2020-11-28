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
	"github.com/lib/pq"
)

var columns = []string{"id", "first_name", "last_name", "email"}

var classColumns = []string{"id", "CreatedAt", "UpdatedAt", "name", "keywords", "categories", "study_buddies"}
var categories = []string{"Math", "Biology"}

func TestClassesInfo(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM courses`)).
		WillReturnRows(sqlmock.NewRows(classColumns).AddRow(1, c, c, "Anatomy", nil, pq.Array(categories), nil))

	obj :=	e.GET("/classes-info").
			Expect().
			Status(http.StatusOK).JSON().Object()

	obj.Keys().Contains("courses")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

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

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(db, w, r)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginUser(db, w, r)
	})

	mux.HandleFunc("/classes-info", func(w http.ResponseWriter, r *http.Request) {
		GetClassesInfo(db, w, r)
	})

	mux.HandleFunc("/updateInvitation", func(w http.ResponseWriter, r *http.Request) {
		UpdateInvitation(db, w, r)
	})

	mux.HandleFunc("/addCourse", func(w http.ResponseWriter, r *http.Request) {
		AddCourse(db, w, r)
	})

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