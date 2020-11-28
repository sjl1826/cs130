package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"regexp"

	"github.com/gavv/httpexpect"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var invitationColumns = []string{"g_id", "CreatedAt", "UpdatedAt", "group_name", "group_id", "receive_id", "receive_name", "type", "status"}

var invRequest = map[string]interface{}{
	"group_name": "Flamingos",
	"group_id": 2,
	"receive_id": 3,
	"receive_name": "Playa",
	"type": true,
}

func TestSendInvitationSuccess(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)	
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	//CreateInvitation
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "invitations"`)).
		WithArgs(AnyTime{}, AnyTime{}, invRequest["group_name"],
				 invRequest["group_id"], invRequest["receive_id"], invRequest["receive_name"], invRequest["type"], false).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "invitations"`)).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))

	//GetInvitation
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "invitations"`)).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "invitations"`)).
		WillReturnRows(sqlmock.NewRows(invitationColumns).AddRow(1, c, c, invRequest["group_name"], invRequest["group_id"],
		invRequest["receive_id"],invRequest["receive_name"], invRequest["type"], false))

	obj :=	e.POST("/sendInvitation").WithJSON(invRequest).
			Expect().
			Status(http.StatusCreated).JSON().Object()

	//Testing the output
	obj.Keys().Contains("group_name", "group_id", "receive_name", "receive_id")
	obj.Value("group_name").Equal(invRequest["group_name"])

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestSendInvitationDatabaseFailure(t *testing.T) {
	mock, DB := GetMock();
	defer DB.Close()

	handler := GetHandler(DB)	
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	//CreateInvitation
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "invitations"`)).
		WithArgs(AnyTime{}, AnyTime{}, invRequest["group_name"],
				 invRequest["group_id"], invRequest["receive_id"], invRequest["receive_name"], invRequest["type"], false).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "invitations"`)).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow("1"))

	//GetInvitation
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "invitations"`)).
		WillReturnError(gorm.ErrRecordNotFound)

	e.POST("/sendInvitation").WithJSON(invRequest).
	Expect().Status(http.StatusNotFound)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}	