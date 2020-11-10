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

// DBMigrate will create and migrate the tables, and then make the relationships
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Token{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Listing{})
	db.AutoMigrate(&Request{})
	return db
}
