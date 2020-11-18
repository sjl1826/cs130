// course.go

package models

import (
	"strconv"
	"time"
	
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Course structure
type Course struct {
	ID          	int 			`gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	Name			string			`json:"name"`
	Keywords		pq.StringArray 	`gorm:"type:varchar(64)[]" json:"keywords"`
	Categories		pq.StringArray 	`gorm:"type:varchar(64)[]" json:"categories"`
	StudyBuddies	pq.Int64Array	`gorm:"type:integer[]" json:"study_buddies"`
	Listings		pq.Int64Array	`gorm:"type:integer[]" json:"listings"`
}

// CreateCourse creates a new course object in database
func (c *Course) CreateCourse(db *gorm.DB) error {
	now := time.Now()
	c.CreatedAt = now
	retVal := db.Create(c).Table("courses").Scan(&c)
	return retVal.Error
}

// UpdateCourse updates course fields
func (c *Course) UpdateCourse(db *gorm.DB) error {
	now := time.Now()
	c.UpdatedAt = now
	retVal := db.Save(&c).Table("courses")
	return retVal.Error
}

// GetCourse retrieves course object from database
func (c *Course) GetCourse(db *gorm.DB) error {
	retVal := db.First(&c, c.ID).Scan(&c).Table("courses")
	return retVal.Error
}

// DeleteCourse deletes course from database
func (c *Course) DeleteCourse(db *gorm.DB) error {
	retVal := db.Exec("DELETE FROM courses WHERE ID=" + strconv.Itoa(c.ID))
	return retVal.Error
}

// AddStudyBuddy adds a new user to the course
func (c *Course) AddStudyBuddy(db *gorm.DB, userID int) error {
	now := time.Now()
	c.UpdatedAt = now
	c.StudyBuddies = append(c.StudyBuddies, int64(userID))
	retVal := db.Save(&c).Table("courses")
	return retVal.Error
}

// RemoveStudyBuddy removes the specified user from the course
func (c *Course) RemoveStudyBuddy(db *gorm.DB, userID int) error {
	now := time.Now()
	c.UpdatedAt = now
	for i, g := range c.StudyBuddies {
		if g == int64(userID) {
			c.StudyBuddies = RemoveElement(c.StudyBuddies, i)
			break
		}
	}
	retVal := db.Save(&c).Table("courses")
	return retVal.Error
}

//Duplicate code here, opportunity to combine :)

// AddListing adds a new listing to the course
func (c *Course) AddListing(db *gorm.DB, listingID int) error {
	now := time.Now()
	c.UpdatedAt = now
	c.Listings = append(c.Listings, int64(listingID))
	retVal := db.Save(&c).Table("courses")
	return retVal.Error
}

// RemoveListing removes the specified listing from the course
func (c *Course) RemoveListing(db *gorm.DB, listingID int) error {
	now := time.Now()
	c.UpdatedAt = now
	for i, g := range c.Listings {
		if g == int64(listingID) {
			c.Listings = RemoveElement(c.Listings, i)
			break
		}
	}
	retVal := db.Save(&c).Table("courses")
	return retVal.Error
}
