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

var userColumns = []string{"u_id", "CreatedAt", "UpdatedAt", 
"first_name", "last_name", "u_email", "password", 
"biography", "discord", "facebook", "timezone", "school_name", "availability"}
func TestRegisterUser(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	//First calls INSERT to Create object. Then SELECT to retrieve it
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs( AnyTime{}, AnyTime{}, "Hunter", "Hunter", "hunter@ymail.com", "", "", "", "", "", "", nil).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows(userColumns).AddRow(1, c, c, "Hunter", "Hunter", "hunter@ymail.com", "", "", "", "", "", "", nil))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE ID=0")).WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	
	
	tempUser := map[string]interface{} {
		"first_name": "Hunter",
		"last_name": "Hunter",
		"u_email": "hunter@ymail.com",
		"password": "",
	}

	e.POST("/register").WithJSON(tempUser).
	Expect().
	Status(http.StatusCreated).JSON().Object().ContainsKey("u_id").ValueEqual("u_id", 0)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLoginUser(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE ID=0")).WillReturnRows(sqlmock.NewRows(userColumns).AddRow(1, c, c, "Hunter", "Hunter", "hunter@ymail.com", "thismypassword", "", "", "", "", "", nil))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT password FROM users where email='hunter@ymail.com'")).WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("thismypassword"))

	e.POST("/login").WithFormField("email", "hunter@ymail.com").WithFormField("password", "thismypassword").
	Expect().
	Status(http.StatusOK).JSON().Object().ContainsKey("access_token")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

var columns = []string{"id", "first_name", "last_name", "email"}


var categoriesCourse = []string{"Physics", "Math"}
var addCourseIDs = []int64{1}
var addCourseRequest = map[string]interface{}{
	"u_id": 1,
	"course_id": 0,
	"course_name": "Physics",
	"keywords": nil,
	"categories": []string{"Physics", "Math"},
}
func TestAddCourse(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE ID=1`)).
		WillReturnRows(sqlmock.NewRows(userColumns).AddRow(1, c, c, "Hunter", "Hunter", "hunter@ymail.com", "gibberish", "", "", "", "", "", nil))
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "courses"`)).
		WithArgs( AnyTime{}, AnyTime{}, "Physics", nil, pq.Array(categoriesCourse), pq.Int64Array(addCourseIDs)).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses"`)).
		WillReturnRows(sqlmock.NewRows(classColumns).AddRow(1, c, c, "Physics", nil, pq.Array(categoriesCourse), nil))


	obj :=	e.PUT("/addCourse").WithJSON(addCourseRequest).
			Expect().
			Status(http.StatusOK).JSON().Object()

	obj.ContainsKey("result").ValueEqual("result", "success")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}


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