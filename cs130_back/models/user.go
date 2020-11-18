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
	Courses			pq.Int64Array	`gorm:"type:integer[]" json:"courses"`
	Groups			pq.Int64Array	`gorm:"type:integer[]" json:"groups"`
	Listings		pq.Int64Array	`gorm:"type:integer[]" json:"listings"`
	Availability	pq.Int64Array	`gorm:"type:integer[]" json:"availability"`
	Invitations		pq.Int64Array	`gorm:"type:integer[]" json:"invitations"`
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

// AddGroup adds a new group to the user
func (u *User) AddGroup(db *gorm.DB, groupID int) error {
	now := time.Now()
	u.UpdatedAt = now
	u.Groups = append(u.Groups, int64(groupID))
	retVal := db.Save(&u).Table("users")
	return retVal.Error
}

// RemoveGroup removes the specified group from the user
func (u *User) RemoveGroup(db *gorm.DB, groupID int) error {
	now := time.Now()
	u.UpdatedAt = now
	for i, g := range u.Groups {
		if g == int64(groupID) {
			u.Groups = RemoveElement(u.Groups, i)
			break
		}
	}
	retVal := db.Save(&u).Table("users")
	return retVal.Error
}

// GetGroups retrieves the group objects under the user
func (u *User) GetGroups(db *gorm.DB, groupList *[]Group) error {
	retVal := db.Raw("SELECT * FROM users WHERE ID=" + strconv.Itoa(u.ID)).Scan(&u)
	for _, g := range u.Groups {
		tempGroup := Group{ID: int(g)}
		db.Raw("SELECT * FROM groups WHERE ID=" + strconv.Itoa(tempGroup.ID)).Scan(&tempGroup)
		(*groupList) = append((*groupList), tempGroup)
	}
	return retVal.Error
}

// AddCourse adds a new course to the user
func (u *User) AddCourse(db *gorm.DB, courseID int) error {
	now := time.Now()
	u.UpdatedAt = now
	u.Courses = append(u.Courses, int64(courseID))
	retVal := db.Save(&u).Table("users")
	return retVal.Error
}

// RemoveCourse removes the specified Course from the user
func (u *User) RemoveCourse(db *gorm.DB, courseID int) error {
	now := time.Now()
	u.UpdatedAt = now
	for i, g := range u.Courses {
		if g == int64(courseID) {
			u.Courses = RemoveElement(u.Courses, i)
			break
		}
	}
	retVal := db.Save(&u).Table("users")
	return retVal.Error
}

// GetCourses retrieves the course objects under the user
func (u *User) GetCourses(db *gorm.DB, courseList *[]Course) error {
	retVal := db.Raw("SELECT * FROM users WHERE ID=" + strconv.Itoa(u.ID)).Scan(&u)
	for _, g := range u.Courses {
		tempCourse := Course{ID: int(g)}
		db.Raw("SELECT * FROM courses WHERE ID=" + strconv.Itoa(tempCourse.ID)).Scan(&tempCourse)
		(*courseList) = append((*courseList), tempCourse)
	}
	return retVal.Error
}

// AddListing adds a new listing to the user
func (u *User) AddListing(db *gorm.DB, listingID int) error {
	now := time.Now()
	u.UpdatedAt = now
	u.Listings = append(u.Listings, int64(listingID))
	retVal := db.Save(&u).Table("users")
	return retVal.Error
}

// RemoveListing removes the specified Listing from the user
func (u *User) RemoveListing(db *gorm.DB, listingID int) error {
	now := time.Now()
	u.UpdatedAt = now
	for i, g := range u.Listings {
		if g == int64(listingID) {
			u.Listings = RemoveElement(u.Listings, i)
			break
		}
	}
	retVal := db.Save(&u).Table("users")
	return retVal.Error
}

// GetListings retrieves the listing objects under the user
func (u *User) GetListings(db *gorm.DB, listingList *[]Listing) error {
	retVal := db.Raw("SELECT * FROM users WHERE ID=" + strconv.Itoa(u.ID)).Scan(&u)
	for _, g := range u.Listings {
		tempListing := Listing{ID: int(g)}
		db.Raw("SELECT * FROM listings WHERE ID=" + strconv.Itoa(tempListing.ID)).Scan(&tempListing)
		(*listingList) = append((*listingList), tempListing)
	}
	return retVal.Error
}

// AddInvitation adds a new invitation to the user
func (u *User) AddInvitation(db *gorm.DB, invitationID int) error {
	now := time.Now()
	u.UpdatedAt = now
	u.Invitations = append(u.Invitations, int64(invitationID))
	retVal := db.Save(&u).Table("users")
	return retVal.Error
}

// RemoveInvitation removes the specified invitation from the user
func (u *User) RemoveInvitation(db *gorm.DB, invitationID int) error {
	now := time.Now()
	u.UpdatedAt = now
	for i, g := range u.Invitations {
		if g == int64(invitationID) {
			u.Invitations = RemoveElement(u.Invitations, i)
			break
		}
	}
	retVal := db.Save(&u).Table("users")
	return retVal.Error
}

// GetInvitations retrieves the invitation objects under the user
func (u *User) GetInvitations(db *gorm.DB, invitationList *[]Invitation) error {
	retVal := db.Raw("SELECT * FROM users WHERE ID=" + strconv.Itoa(u.ID)).Scan(&u)
	for _, g := range u.Invitations {
		tempInvitation := Invitation{ID: int(g)}
		db.Raw("SELECT * FROM invitations WHERE ID=" + strconv.Itoa(tempInvitation.ID)).Scan(&tempInvitation)
		(*invitationList) = append((*invitationList), tempInvitation)
	}
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
