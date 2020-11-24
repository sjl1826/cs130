//seed.go

package seeds

import (
	"cs130_back/models"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
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
	}

	for i := range courses {
		if err := courses[i].CreateCourse(db); err != nil {  	//Courses
			log.Fatal("Error seeding courses", err)
		}
	}

	for i := range groups { 
		if err := groups[i].CreateGroup(db); err != nil {		//Groups
			log.Fatal("Error seeding groups", err)
		}
	}

	for i := range listings {
		if err := listings[i].CreateListing(db); err != nil {	//Listings
			log.Fatal("Error seeding listings", err)
		}
	}

	for i := range invitations {
		if err := invitations[i].CreateInvitation(db); err != nil {	//Invitations
			log.Fatal("Error seeding invitations", err)
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
		StudyBuddies:	pq.Int64Array {3},
	},
	models.Course{
		ID: 		155,	
		Name: 		"Human Anatomy",
		Keywords: 	[]string{"Biology", "Reproductive System", "Skeletal System"},
		Categories: []string{"College", "Physical Sciences", "Human Anatomy"},
		StudyBuddies:	pq.Int64Array {3},
	},
	models.Course{
		ID: 		117,
		Name: 		"Linear Algebra",
		Keywords: 	[]string{"Math", "Graphs", "Matrices"},
		Categories: []string{"College", "Mathematics", "Linear Algebra"}, 
	},
	models.Course{
		ID: 		156,	
		Name: 		"Beginner Psychology",
		Keywords: 	[]string{"Emotions", "Psychopathy", "Disorders"},
		Categories: []string{"College", "Social Sciences", "Beginner Psychology"}, 
	},
	models.Course{
		ID: 		157,
		Name: 		"Calculus 1",
		Keywords: 	[]string{"Derivatives", "Integrals", "Limits"},
		Categories: []string{"College", "Mathematics", "Calculus 1"}, 
	},
	models.Course{
		ID: 		158,	
		Name: 		"AP Calculus AB",
		Keywords: 	[]string{"Derivatives", "Integrals", "Limits"},
		Categories: []string{"High School", "Mathematics", "AP Calculus AB"}, 
	},
}

var groups = []models.Group{
	models.Group{
		Name: 		"48ers",
		CourseName: "Human Anatomy",
		CourseID: 	155,
		AdminID: 	1,
		Members:	pq.Int64Array {3},
	},
	models.Group{
		Name: 		"Hedrick Homies",
		CourseName: "Lower Division Linear Algebra",
		CourseID: 	115,
		AdminID: 	2,
		Members:	pq.Int64Array {3},
	},
}

var listings = []models.Listing{
	models.Listing{
		ID:			99,
		CourseName: "Human Anatomy",
		Poster: 	3,
		CourseID: 	155,
		Tags:		[]string {"3.9+ GPA only", "No public school kids"},
	},
	models.Listing{
		ID:			100,
		CourseName: "Human Anatomy",
		Poster: 	3,
		CourseID: 	155,
		Tags:		[]string {"Casual"},
	},
}


var invitations = []models.Invitation{
	models.Invitation{
		GroupName: 	"48ers",
		GroupID: 	1,
		ReceiveID: 	3,
		ReceiveName: "Edgar Garcia",
		Type: 		false,
		Status:		false,
	},
	models.Invitation{
		GroupName: 	"Hedrick Homies",
		GroupID: 	2,
		ReceiveID: 	3,
		ReceiveName: "Edgar Garcias",
		Type: 		false,
		Status:		false,
	},
	models.Invitation{
		GroupName: 	"48ers",
		GroupID: 	1,
		ReceiveID: 	4,
		ReceiveName: "Jerry Roach",
		Type: 		true,
		Status:		false,
	},
}