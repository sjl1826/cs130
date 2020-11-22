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

func (g *Group) AddMember(db *gorm.DB, userID int) error {
	now := time.Now()
	g.UpdatedAt = now
	g.Members = append(g.Members, int64(userID))
	retVal := db.Save(&g).Table("groups")
	return retVal.Error
}

// GetAvailability retrieves the availability object of the group

// func (g *Group) GetAvailability(db *gorm.DB, availability *[]int64) error {
// 	retVal := db.Raw("SELECT * FROM users WHERE ID=" + strconv.Itoa(u.ID)).Scan(&u)
// 	for _, g := range u.Courses {
// 		tempCourse := Course{ID: int(g)}
// 		db.Raw("SELECT * FROM courses WHERE ID=" + strconv.Itoa(tempCourse.ID)).Scan(&tempCourse)
// 		(*courseList) = append((*courseList), tempCourse)
// 	}
// 	return retVal.Error

	
// }


func computeOverlap(avSet [][]int64) []int64 {
	var overlap []int64

	for i := 0; i < len(avSet); i++ {
        for j := 0; j < len(avSet[0]); j++ {
            overlap[j] = overlap[j] + avSet[i][j]
        }
    }

	return overlap
}
