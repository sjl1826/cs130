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
	// TODO: studybuddies and listings
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
