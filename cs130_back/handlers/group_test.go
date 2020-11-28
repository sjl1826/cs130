package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"regexp"

	"github.com/gavv/httpexpect"
	"github.com/DATA-DOG/go-sqlmock"
)

var groupColumns = []string{"created_at", "updated_at", "name", "course_name", "course_id", "admin_id", "members"}

var grpRequest = map[string]interface{}{
	"admin_id": 3,
	"name": "Frozen Flamingos",
	"course_id": 155,
}

func TestCreateGroupSuccess(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)	
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	//CreateGroup
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "groups"`)).
		WithArgs(AnyTime{}, AnyTime{}, grpRequest["name"], "", grpRequest["course_id"], grpRequest["admin_id"], nil).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "groups"`)).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))

	//GetGroup
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "groups"`)).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "groups"`)).
		WillReturnRows(sqlmock.NewRows(groupColumns).AddRow(c, c, grpRequest["name"], nil, grpRequest["course_id"], grpRequest["admin_id"], nil))

	obj :=	e.POST("/createGroup").WithJSON(grpRequest).
			Expect().
			Status(http.StatusCreated).JSON().Object()

	//Testing the output
	obj.Keys().Contains("name")
	obj.Value("name").Equal(grpRequest["name"])

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateGroupBadRequest(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)	
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	badGrpRequest := map[string]interface{}{
		"admin_id": 3,
		"name": 1,	//integer in place of string
		"course_id": 155, 
	}
	e.POST("/createGroup").WithJSON(badGrpRequest).
	Expect().
	Status(http.StatusBadRequest)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}