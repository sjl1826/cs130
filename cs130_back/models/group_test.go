package models

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var t = time.Now()

var groupColumns = []string{"created_at", "updated_at", "name", "course_name", "course_id", "admin_id", "members"}

var g = Group{
	ID:         6,
	Name:       "Studdie4life",
	CourseName: "Human Anatomy",
	CourseID:   155,
	AdminID:    1,
}

func TestUpdateGroup(t *testing.T) {

	mock, DB := GetMock()
	defer DB.Close()

	//Calls Update on object to update it, then SELECT to retrieve it
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").
		WithArgs(AnyTime{}, g.Name, g.CourseName, g.CourseID, g.AdminID, g.Members, g.ID).WillReturnResult(sqlmock.NewResult(0, 0))
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

func TestGetMeetingTime(t *testing.T) {
	var meetingTime string
	mock, DB := GetMock()
	defer DB.Close()

	var members = []int64{1, 2}
	availability := make([]int64, 336)
	availability[0] = 1

	// Retrieve group
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM groups`)).
		WillReturnRows(sqlmock.NewRows(groupColumns).AddRow(c, c, g.Name, g.CourseName, g.CourseID, g.AdminID, pq.Array(members)))
	// Get availability of users in group
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).
		WillReturnRows(sqlmock.NewRows(userColumns).AddRow(1, c, c, "Aaron", "Pozer", "ap@gmail.com", "", "", "", "", "", "", pq.Array(availability)))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).
		WillReturnRows(sqlmock.NewRows(userColumns).AddRow(2, c, c, "Blake", "Dozer", "bd@gmail.com", "", "", "", "", "", "", pq.Array(availability)))

	// now we execute our method
	if err := g.GetMeetingTime(DB, &meetingTime); err != nil {
		t.Errorf("error was not expected while updating group: %s", err)
	}

	result := "The group meeting time is Monday at 12:00 am"
	assert.Equal(t, meetingTime, result, "Both strings must be equal") 

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
