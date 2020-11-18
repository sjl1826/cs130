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
	Institution     string          `json:"institution"`
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

// GetCategories retrieves all categories that exist in the database
func (c *Course) GetCategories(db *gorm.DB, categories *[]string, institution string) error {
	// retVal := db.Raw("SELECT DISTINCT * FROM (SELECT ARRAY_AGG(val) FROM (SELECT UNNEST(categories) FROM courses WHERE institution = " + institution + " AS val))").Scan(&categories)
	retVal := db.Raw("SELECT ARRAY_AGG(DISTINCT c) FROM (SELECT categories FROM courses WHERE institution LIKE " + "'"+institution+"' AS c) AS u").Scan(&categories)
	return retVal.Error
}

// Getsubcategories retrieves all subcategories that exist in the database for a category and an institution
func (c *Course) GetSubcategories(db *gorm.DB, subcategories *[]string, category string, institution string) error {
	retVal := db.Raw("SELECT name FROM courses WHERE institution = " + "'" + institution + "'" + " AND " + category + " = ANY(categories)").Scan(&subcategories)
	return retVal.Error
}

// GetCoursesByCategory retrieves all course objects with the category specified from database
func (c *Course) GetCoursesByCategory(db *gorm.DB, courseList *[]Course, category string, institution string) error {
	retVal := db.Raw("SELECT * FROM courses WHERE " + category + " = ANY (categories) AND " + institution + " = institution").Scan(&courseList)
	return retVal.Error
}

// GetCoursesBySubcategory retrieves all course objects with the subcategory specified from database
func (c *Course) GetCoursesBySubcategory(db *gorm.DB, courseList *[]Course, keywords *[]string, subcategory string, institution string) error {
	retVal := db.Raw("SELECT * FROM courses WHERE " + subcategory + " = name AND " + institution + " = institution").Scan(&courseList)
	db.Raw("SELECT keywords FROM courses WHERE " + subcategory + " = name AND " + institution + " = institution LIMIT 1").Scan(&keywords)
	return retVal.Error
}