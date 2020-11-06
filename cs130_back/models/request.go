// request.go

package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// Request structure
type Request struct {
	ID          	int 			`gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	GroupName		string			`json:"group_name"`
	GroupID			int				`json:"group_id"`
	ReceiveID		int				`json:"receive_id`
	Type			bool			`json:"type"`
	Status			bool			`json:"status"`
}

// CreateRequest creates a new request object in database
func (r *Request) CreateRequest(db *gorm.DB) error {
	now := time.Now()
	r.CreatedAt = now
	retVal := db.Create(r).Table("request").Scan(&r)
	return retVal.Error
}

// UpdateRequest updates request fields
func (r *Request) UpdateRequest(db *gorm.DB) error {
	now := time.Now()
	r.UpdatedAt = now
	retVal := db.Save(&r).Table("requests")
	return retVal.Error
}

// GetRequest retrieves request object from database
func (r *Request) GetRequest(db *gorm.DB) error {
	retVal := db.First(&r, r.ID).Scan(&r).Table("requests")
	return retVal.Error
}

// DeleteRequest deletes request from database
func (r *Request) DeleteRequest(db *gorm.DB) error {
	retVal := db.Exec("DELETE FROM erquests WHERE ID=" + strconv.Itoa(r.ID))
	return retVal.Error
}
