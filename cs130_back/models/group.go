// group.go

package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Group structure
type Group struct {
	ID          int `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string        `json:"name"`
	CourseName  string        `json:"course_name"`
	CourseID    int           `json:"course_id"`
	AdminID     int           `json:"admin_id"`
	Members     pq.Int64Array `gorm:"type:integer[]" json:"members"`
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

// AddMember adds a new user to the group
func (g *Group) AddMember(db *gorm.DB, userID int) error {
	now := time.Now()
	g.UpdatedAt = now
	g.Members = append(g.Members, int64(userID))
	retVal := db.Save(&g).Table("groups")
	return retVal.Error
}

// RemoveMember removes the specified user from the group
func (g *Group) RemoveMember(db *gorm.DB, userID int) error {
	now := time.Now()
	g.UpdatedAt = now
	for i, j := range g.Members {
		if j == int64(userID) {
			g.Members = RemoveElement(g.Members, i)
			break
		}
	}
	retVal := db.Save(&g).Table("groups")
	return retVal.Error
}

// GetMembers returns the groups members
func (g *Group) GetMembers(db *gorm.DB, members *[]User) error {
	retVal := db.Raw("SELECT * FROM groups WHERE ID=" + strconv.Itoa(g.ID)).Scan(&g)
	for _, j := range g.Members {
		tempMember := User{ID: int(j)}
		db.Raw("SELECT * FROM users WHERE ID=" + strconv.Itoa(tempMember.ID)).Scan(&tempMember)
		(*members) = append((*members), tempMember)
	}
	return retVal.Error
}

const numSlots = 336

//GetAvailability retrieves the availability object of the group
func (g *Group) GetAvailability(db *gorm.DB, availability *[numSlots]int64) error {
	var avSet *[][numSlots]int64
	retVal := db.Raw("SELECT * FROM groups WHERE ID=" + strconv.Itoa(g.ID)).Scan(&g)
	for _, j := range g.Members {
		tempMember := User{ID: int(j)}
		db.Raw("SELECT * FROM users WHERE ID=" + strconv.Itoa(tempMember.ID)).Scan(&tempMember)
		var tempAvailability [numSlots]int64
		tempMember.Availability.Scan(&tempAvailability)
		(*avSet) = append((*avSet), tempAvailability)

	}

	availability = computeOverlap(avSet)

	return retVal.Error
}

func computeOverlap(avSet *[][numSlots]int64) *[numSlots]int64 {
	var overlap [numSlots]int64

	for _, i := range *avSet {
		for j := 0; j < len(i); j++ {
			overlap[j] = overlap[j] + i[j]
		}
	}

	return &overlap
}

//GetMeetingTime retrieves the meeting time string for the group
func (g *Group) GetMeetingTime(db *gorm.DB, m *string) error {
	var availability *[numSlots]int64

	retVal := g.GetAvailability(db, availability)

	var highest int64 = 0

	var highestLoc int = -1

	for loc, val := range *availability {
		if val > highest {
			highest = val
			highestLoc = loc
		}
	}

	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	day := highestLoc / 48 // 335 /48 = 6 = Saturday

	timeBlock := highestLoc % 48 //335 mod 48 = 47(11:30 PM)

	hour := (timeBlock / 2) % 12 // 47/2 = 23 mod 12 = 11

	if hour == 0 {
		hour = 12
	}

	var minutes string

	if timeBlock%2 == 1 { //  odd parity
		minutes = ":30"
	} else {
		minutes = ":00"
	}

	var meridie string

	if timeBlock > 23 {
		meridie = "pm"
	} else {
		meridie = "am"
	}

	*m = fmt.Sprintf("The Group Meeting Day is %s at %d%s %s", days[day], hour, minutes, meridie)

	return retVal
}
