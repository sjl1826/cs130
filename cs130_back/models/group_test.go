package models

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var t = time.Now()

var groupColumns = []string{"created_at", "updated_at", "name", "course_name", "course_id", "admin_id", "meeting_time", "members"}


func TestUpdateGroup(t *testing.T) {
	var g = Group {
		ID:			6,
		Name: 		"Studdie4life",
		CourseName: "Human Anatomy",
		CourseID: 	155,
		AdminID: 	1,
		MeetingTime: "4:21PM",
	}

	mock, DB := GetMock()
	defer DB.Close()

	//Calls Update on object to update it, then SELECT to retrieve it
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").
		WithArgs(AnyTime{}, g.Name, g.CourseName, g.CourseID, g.AdminID, g.MeetingTime, g.Members, g.ID).WillReturnResult(sqlmock.NewResult(0,0))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "groups" WHERE "groups"."id" = $1 ORDER BY "groups"."id" ASC LIMIT 1`)).
		WithArgs(g.ID).WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))

	// now we execute our method
	if err := g.UpdateGroup(DB); err != nil {
		t.Errorf("error was not expected while updating group: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}