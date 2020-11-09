// group.go

package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Group structure
type Group struct {
	ID          	int 			`gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	Name			string			`json:"name"`
	CourseName		string			`json:"course_name"`
	CourseID		int				`json:"course_id"`
	AdminID			int				`json:"admin_id"`
	MeetingTime		string			`json:"meeting_time"`
	Members			pq.Int64Array	`gorm:"type:integer[]" json:"members"`
	Requests		pq.Int64Array	`gorm:"type:integer[]" json:"requests"`
}

// CreateGroup creates a new group object in database
func (g *Group) CreateGroup(db *gorm.DB) error {
	now := time.Now()
	g.CreatedAt = now
	retVal := db.Create(g).Table("groups").Scan(&g)
	return retVal.Error
}

// UpdateGroup updates group fields
func (g *Group) UpdateGroup(db *gorm.DB) error {
	now := time.Now()
	g.UpdatedAt = now
	retVal := db.Save(&g).Table("groups")
	return retVal.Error
}

// GetGroup retrieves group object from database
func (g *Group) GetGroup(db *gorm.DB) error {
	retVal := db.First(&g, g.ID).Scan(&g).Table("groups")
	return retVal.Error
}

// DeleteGroup deletes group from database
func (g *Group) DeleteGroup(db *gorm.DB) error {
	retVal := db.Exec("DELETE FROM groups WHERE ID=" + strconv.Itoa(g.ID))
	return retVal.Error
}
