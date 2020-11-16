// listing.go

package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Listing structure
type Listing struct {
	ID          	int 			`gorm:"primary_key;auto_increment" json:"listing_id"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	CourseName		string			`json:"course_name"`
	Poster			int				`json:"poster"`
	CourseID		int				`json:"course_id"`
	Description		string			`json:"text_description"`
	GroupID			int				`json:"group_id"`		//optional group		
	Tags			pq.StringArray 	`gorm:"type:varchar(64)[]" json:"tags"`
}

// CreateListing creates a new listing object in database
func (l *Listing) CreateListing(db *gorm.DB) error {
	now := time.Now()
	l.CreatedAt = now
	retVal := db.Create(l).Table("listings").Scan(&l)
	return retVal.Error
}

// UpdateListing updates listing fields
func (l *Listing) UpdateListing(db *gorm.DB) error {
	now := time.Now()
	l.UpdatedAt = now
	retVal := db.Save(&l).Table("listings")
	return retVal.Error
}

// GetListing retrieves listing object from database
func (l *Listing) GetListing(db *gorm.DB) error {
	retVal := db.First(&l, l.ID).Scan(&l).Table("listings")
	return retVal.Error
}

// DeleteListing deletes listing from database
func (l *Listing) DeleteListing(db *gorm.DB) error {
	retVal := db.Exec("DELETE FROM listings WHERE ID=" + strconv.Itoa(l.ID))
	return retVal.Error
}
