// subject.go

package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// Subject structure
type Subject struct {
	ID          	int 	`gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	Description		string	`json:"description"`
	Title			string 	`json:"title"`
}

// CreateSubject creates a new subject object in database
func (s *Subject) CreateSubject(db *gorm.DB) error {
	now := time.Now()
	s.CreatedAt = now
	retVal := db.Create(s).Table("subjects").Scan(&s)
	return retVal.Error
}

// UpdateSubject updates subject fields
func (s *Subject) UpdateSubject(db *gorm.DB) error {
	now := time.Now()
	s.UpdatedAt = now
	retVal := db.Save(&s).Table("subjects")
	return retVal.Error
}

// GetSubject retrieves subject object from database
func (s *Subject) GetSubject(db *gorm.DB) error {
	retVal := db.First(&s, s.ID).Scan(&s).Table("subjects")
	return retVal.Error
}

// DeleteSubject deletes subject from database
func (s *Subject) DeleteSubject(db *gorm.DB) error {
	retVal := db.Exec("DELETE FROM subjects WHERE ID=" + strconv.Itoa(s.ID))
	return retVal.Error
}
