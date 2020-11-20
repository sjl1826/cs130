package models

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var n = time.Now()

var userColumns = []string{"id", "created_at", "updated_at", "first_name", "last_name", "email", "password", "biography", "discord", "facebook", "timezone", "school_name", "availability"}

func TestCreateUser(t *testing.T) {
	var u = User{
		FirstName: "Hunter",
		LastName: "Greece",
		Email:	"hgtrece@gmail.com",
		Timezone: "PST",
	}

	mock, DB := GetMock()
	defer DB.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs(AnyTime{}, AnyTime{}, u.FirstName, u.LastName, u.Email, u.Password, u.Biography, u.Discord, u.Facebook, u.Timezone, u.SchoolName, u.Availability).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows(userColumns).AddRow(1, c, c, u.FirstName, u.LastName, u.Email, u.Password, u.Biography, u.Discord, u.Facebook, u.Timezone, u.SchoolName, u.Availability))

	// now we execute our method
	if err := u.CreateUser(DB); err != nil {
		t.Errorf("error was not expected while creating user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}