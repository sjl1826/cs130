// user.go

package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// RemoveElement moves the end of a slice to the index and returns a new slice (does not preserve order)
func RemoveElement(s []int64, i int) []int64 {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// User holds all fields for registered users
type User struct {
	ID				int		`gorm:"primary_key;auto_increment" json:"u_id"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	FirstName		string	`json:"first_name"`
	LastName		string	`json:"last_name"`
	Email			string	`gorm:"unique" json:"u_email"`
	Password		string	`json:"password"`
	Biography		string	`json:"biography"`
	Discord			string 	`json:"discord"`
	Facebook		string 	`json:"facebook"`
	Timezone		string 	`json:"timezone"`
	SchoolName		string 	`json:"school_name"` 
	Availability	pq.Int64Array	`gorm:"type:integer[]" json:"availability"`
}

// CreateUser creates the user specified
func (u *User) CreateUser(db *gorm.DB) error {
	now := time.Now()
	u.CreatedAt = now
	retVal := db.Create(u).Table("users").Scan(&u)
	return retVal.Error
}

// GetUser gets specified user and returns it
func (u *User) GetUser(db *gorm.DB) error {
	retVal := db.Raw("SELECT * FROM users WHERE ID=" + strconv.Itoa(u.ID)).Scan(&u)
	return retVal.Error
}

// UpdateUser updates user fields
func (u *User) UpdateUser(db *gorm.DB) error {
	now := time.Now()
	u.UpdatedAt = now
	retVal := db.Save(&u).Table("users")
	return retVal.Error
}

// DeleteUser deletes the specified user
func (u *User) DeleteUser(db *gorm.DB) error {
	retVal := db.Exec("DELETE FROM users WHERE ID=" + strconv.Itoa(u.ID))
	return retVal.Error
}

// GetByEmail gets the user by email, not ID
func (u *User) GetByEmail(db *gorm.DB) error {
	row := db.Raw("SELECT id FROM users WHERE email='" + u.Email + "'").Row()
	var i int
	row.Scan(&i)
	u.ID = i
	error := u.GetUser(db)
	return error
}

// GetPassword gets the hashed password for the user
func (u *User) GetPassword(db *gorm.DB) {
	row := db.Raw("SELECT password FROM users where email='" + u.Email + "'").Row()
	var pass string
	row.Scan(&pass)
	u.Password = pass
}

// GetGroups retrieves the group objects under the user
func (u *User) GetGroups(db *gorm.DB, groupList *[]Group) error {
	retVal := db.Raw("SELECT * FROM groups WHERE " + strconv.Itoa(u.ID) + " = ANY (members)").Scan(&groupList)
	return retVal.Error
}

// GetCourses retrieves the course objects under the user
func (u *User) GetCourses(db *gorm.DB, courseList *[]Course) error {
	retVal := db.Raw("SELECT * FROM courses WHERE " + strconv.Itoa(u.ID) + " = ANY (study_buddies)").Scan(&courseList)
	return retVal.Error
}

// GetListings retrieves the listing objects under the user
func (u *User) GetListings(db *gorm.DB, listingList *[]Listing) error {
	retVal := db.Raw("SELECT * FROM listings WHERE " + strconv.Itoa(u.ID) + " = poster").Scan(&listingList)
	return retVal.Error
}

// GetInvitations retrieves the invitation objects under the user
func (u *User) GetInvitations(db *gorm.DB, invitationList *[]Invitation) error {
	retVal := db.Raw("SELECT * FROM invitations WHERE " + strconv.Itoa(u.ID) + " = receive_id").Scan(&invitationList)
	return retVal.Error
}

// DBMigrate will create and migrate the tables, and then make the relationships
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Token{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Listing{})
	db.AutoMigrate(&Invitation{})
	return db
}
