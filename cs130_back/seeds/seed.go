//seed.go

package seeds

import (
	"cs130_back/models"
	"log"

	"github.com/jinzhu/gorm"
)

//Migrates tables and plants seed data
func Seed (db *gorm.DB) *gorm.DB {

	//Drops all tables
	err := db.DropTableIfExists(
		&models.User{},
		&models.Course{},
		&models.Token{},
		&models.Group{},
		&models.Listing{},
		&models.Invitation{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	//Migrates all tables
	err = db.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Token{},
		&models.Group{},
		&models.Listing{},
		&models.Invitation{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	//Seeding all seed objects
	for i := range users {
		if err := users[i].CreateUser(db); err != nil {   		//Users
			log.Fatal("Error seeding users", err)
		}

		if err := courses[i].CreateCourse(db); err != nil {  	//Courses
			log.Fatal("Error seeding courses", err)
		}

		if err := groups[i].CreateGroup(db); err != nil {		//Groups
			log.Fatal("Error seeding groups", err)
		}

		if err := listings[i].CreateListing(db); err != nil {	//Listings
			log.Fatal("Error seeding listings", err)
		}
	}

	return db;
}

//Seed Objects

var users = []models.User{
	models.User{
		FirstName: 	"Hunter",
		LastName: 	"Hunter",
		Email: 		"hunter@ymail.com",
		Password: 	"hunter2",
	},
	models.User{
		FirstName: 	"Blake",
		LastName: 	"Bradley",
		Email: 		"brad@ymail.com",
		Password: 	"hunter3",
	},
}

var courses = []models.Course{
	models.Course{
		ID: 		115,
		Name: 		"Lower Division Linear Algebra",
		Keywords: 	[]string{"Math", "Graphs", "Matrices"},
		Categories: []string{"College", "Mathematics", "Linear Algebra"}, 
	},
	models.Course{
		ID: 		155,	
		Name: 		"Human Anatomy",
		Keywords: 	[]string{"Biology", "Reproductive System", "Skeletal System"},
		Categories: []string{"College", "Physical Sciences", "Human Anatomy"}, 
		Listings:	[]int64{99, 100},
	},
}

var groups = []models.Group{
	models.Group{
		Name: 		"48ers",
		CourseName: "Human Anatomy",
		CourseID: 	155,
		AdminID: 	1,
	},
	models.Group{
		Name: 		"Hedrick Homies",
		CourseName: "Lower Division Linear Algebra",
		CourseID: 	115,
		AdminID: 	2,
	},
}

var listings = []models.Listing{
	models.Listing{
		ID:			99,
		CourseName: "Human Anatomy",
		Poster: 	1,
		CourseID: 	155,
		Keywords: 	[]string{"Biology", "Reproductive System", "Skeletal System"},
		Tags:		[]string {"3.9+ GPA only", "No public school kids"},
	},
	models.Listing{
		ID:			100,
		CourseName: "Human Anatomy",
		Poster: 	1,
		CourseID: 	155,
		Keywords: 	[]string{"Biology", "Reproductive System", "Skeletal System"},
		Tags:		[]string {"Casual"},
	},
}