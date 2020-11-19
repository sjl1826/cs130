package models

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

var c = time.Now()

var listingColumns = []string{"created_at", "updated_at", "course_name", "poster", "course_id", "description", "group_id", "tags"}

type AnyTime struct{}
// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

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

func TestCreateListing(t *testing.T) {
	var l = Listing{
		CourseName:  	"Human Anatomy",
		Poster:     	 1,
		CourseID:    	155,
		Description: 	"Testing",
	}

	mock, DB := GetMock();
	defer DB.Close()

	//First calls INSERT to Create object. Then SELECT to retrieve it
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "listings"`)).
		WithArgs(AnyTime{}, AnyTime{}, "Human Anatomy", 1, 155, "Testing", 0, nil).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "listings"`)).
		WillReturnRows(sqlmock.NewRows(listingColumns).AddRow(c, c, l.CourseName, l.Poster, l.CourseID, "hi", 0, nil)) //only thing that works

	// now we execute our method
	if err := l.CreateListing(DB); err != nil {
		t.Errorf("error was not expected while creating listing: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateListing(t *testing.T) {
	var l = Listing{
		ID:				1,
		CourseName:  	"Human Anatomy",
		Poster:     	 1,
		CourseID:    	155,
		Description: 	"Testing",
	}

	mock, DB := GetMock();
	defer DB.Close()

	//Calls Update on object to update it, then SELECT to retrieve it
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").
		WithArgs(AnyTime{}, "Human Anatomy", 1, 155, "Testing", 0, nil, 1).WillReturnResult(sqlmock.NewResult(0,0))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "listings" WHERE "listings"."id" = $1 ORDER BY "listings"."id" ASC LIMIT 1`)).
		WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	

	// now we execute our method
	if err := l.UpdateListing(DB); err != nil {
		t.Errorf("error was not expected while updating listing: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetListing(t *testing.T) {
	var l = Listing{
		ID:				1,
		CourseName:  	"Human Anatomy",
		Poster:     	 1,
		CourseID:    	155,
		Description: 	"Testing",
	}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //mock sql db
	require.NoError(t, err)

	DB, err := gorm.Open("postgres", db) //open gorm db
	require.NoError(t, err)
	defer DB.Close()

	//Calls SELECT to retrieve object
	mock.ExpectQuery(`SELECT * FROM "listings"  WHERE "listings"."id" = $1 AND (("listings"."id" = 1)) ORDER BY "listings"."id" ASC LIMIT 1`).
		WithArgs(l.ID).WillReturnRows(sqlmock.NewRows(listingColumns).AddRow(c, c, l.CourseName, l.Poster, l.CourseID, "hi", 0, nil))

	mock.ExpectQuery(`SELECT * FROM "listings"  WHERE "listings"."id" = $1 ORDER BY "listings"."id" ASC`).
		WithArgs(l.ID).WillReturnRows(sqlmock.NewRows(listingColumns).AddRow(c, c, l.CourseName, l.Poster, l.CourseID, "hi", 0, nil))
	 	
	// now we execute our method
	if err := l.GetListing(DB); err != nil {
		t.Errorf("error was not expected while getting listing: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteListing(t *testing.T) {
	var l = Listing{
		ID:				1,
		CourseName:  	"Human Anatomy",
		Poster:     	 1,
		CourseID:    	155,
		Description: 	"Testing",
	}
	mock, DB := GetMock();
	defer DB.Close()

	//Calls DELETE to delete object
	mock.ExpectExec("DELETE FROM listings WHERE ID=1" ).WillReturnResult(sqlmock.NewResult(0, 0))

	// now we execute our method
	if err := l.DeleteListing(DB); err != nil {
		t.Errorf("error was not expected while deleting listing: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
