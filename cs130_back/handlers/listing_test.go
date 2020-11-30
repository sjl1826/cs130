package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"regexp"
	"cs130_back/models"

	"github.com/gavv/httpexpect"
	"github.com/DATA-DOG/go-sqlmock"
)

var c = time.Now()
var listingColumns = []string{"created_at", "updated_at", "course_name", "poster", "course_id", "description", "group_id", "tags"}

func TestCreateListingSuccess(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)	
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	var l = models.Listing{
		Poster:     	 1,
		CourseID:    	0,
	}

	//First calls INSERT to Create object. Then SELECT to retrieve it
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "listings"`)).
		WithArgs(AnyTime{}, AnyTime{}, l.CourseName, l.Poster, l.CourseID, l.Description, l.GroupID, l.GroupName, l.Tags).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "listings"`)).
		WillReturnRows(sqlmock.NewRows(listingColumns).AddRow(c, c, "TEST COURSE", l.Poster, l.CourseID, l.Description, l.GroupID, l.Tags)) 
	
	//GetListing
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "listings"`)).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "listings"`)).
	WillReturnRows(sqlmock.NewRows(listingColumns).AddRow(c, c, "TEST COURSE", l.Poster, l.CourseID, l.Description, l.GroupID, l.Tags))
	
	list := map[string]interface{}{
		"Poster": 1,
		"CourseID": 2,
		"Description": "hello",
		"Tags": nil,
	}

	e.POST("/createListing").WithJSON(list).
	Expect().
	Status(http.StatusCreated).JSON().Object().ContainsKey("course_name").ValueEqual("course_name", "TEST COURSE")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
