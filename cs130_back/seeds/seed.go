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

func setAvailability(option bool) []int64 {
	availability1 := make([]int64, 336)
	availability1[0] = 1
	availability1[1] = 1
	availability2 := make([]int64, 336)
	availability2[1] = 1
	if option {
		return availability1
	}
	return availability2
}

//Seed Objects

var users = []models.User{
	models.User{
		FirstName: 	"Hunter",
		LastName: 	"Cook",
		Email: 		"hunter@gmail.com",
		Password: 	"hunter2",
		SchoolName: "UCLA",
		Availability: setAvailability(true),
	},
	models.User{
		FirstName: 	"Blake",
		LastName: 	"Bradley",
		Email: 		"brad@ymail.com",
		Password: 	"hunter3",
		SchoolName: "USC",
		Availability: setAvailability(false),
	},
	models.User{
		FirstName: 	"London",
		LastName: 	"Tipton",
		Email: 		"bostonsocialite@gmail.com",
		Password: 	"hunter3",
		Availability: setAvailability(true),
	},
	models.User{
		FirstName: 	"Ash",
		LastName: 	"Ketchum",
		Email: 		"gonnacatchemall@gmail.com",
		Password: 	"hunter3",
		Availability: setAvailability(false),
	},
	models.User{
		FirstName: 	"Meredith",
		LastName: 	"Gray",
		Email: 		"surgeryjunkie@gmail.com",
		Password: 	"hunter3",
		SchoolName: "University of Washington",
		Availability: setAvailability(true),
	},
	models.User{
		FirstName: 	"Derek",
		LastName: 	"Shepard",
		Email: 		"bigbrain@gmail.com",
		Discord: "derek#1255",
		Password: 	"hunter3",
		SchoolName: "University of Washington",
		Availability: setAvailability(true),
	},
	models.User{
		FirstName: 	"Thomas",
		LastName: 	"Brady",
		Email: 		"footballthrow@gmail.com",
		Discord: "tom#1257",
		Password: 	"hunter3",
		SchoolName: "LSU",
		Availability: setAvailability(true),
	},
}

var courses = []models.Course{
	models.Course{
		ID: 		115,
		Name: 		"Lower Division Linear Algebra",
		Keywords: 	[]string{
			"Factorization", "Transpose Matrix", "Fourier Transform", "Identity Matrix",
			"Independent Vectors", "Lucas Numbers", "Traces", "Nullspaces",
		},
		Categories: []string{"College", "Mathematics", "Lower Division Linear Algebra"}, 
		StudyBuddies:	pq.Int64Array {1,2,4},
	},
	models.Course{
		ID: 		116,
		Name: 		"Introduction to Probability",
		Keywords: 	[]string{
			"Discrete random variables", "Discrete Probability Distributions", "Continuous probability distributions", "Central Limit Theorem",
			"Normal Distribution",
		},
		Categories: []string{"College", "Statistics", "Introduction to Probability"}, 
		StudyBuddies:	pq.Int64Array {1,2,4},
	},
	models.Course{
		ID: 		119,
		Name: 		"Statistical Programming with R",
		Keywords: 	[]string{
			"R Programming", "Statistical Programming",
		},
		Categories: []string{"College", "Statistics", "Statistical Programming with R"}, 
	},
	models.Course{
		ID: 		155,	
		Name: 		"Human Anatomy",
		Keywords: 	[]string{"Organ Systems", "Reproductive System", "Skeletal System", "Movements", "Anatomical Regions",},
		Categories: []string{"College", "Life Sciences", "Human Anatomy"},
		StudyBuddies:	pq.Int64Array {1,2,5,6},
	},
	models.Course{
		ID: 		117,
		Name: 		"Linear Algebra",
		Keywords: 	[]string{
			"Factorization", "Transpose Matrix", "Fourier Transform", "Identity Matrix",
			"Independent Vectors", "Lucas Numbers", "Traces", "Nullspaces",
		},
		Categories: []string{"College", "Mathematics", "Linear Algebra"}, 
		StudyBuddies:	pq.Int64Array {5,6},
	},
	models.Course{
		ID: 		156,	
		Name: 		"Beginner Psychology",
		Keywords: 	[]string{"Functionalism", "Psychopathy", "Disorders", "Behaviors", "Abnormal behavior", "Environment", "Hereditary",},
		Categories: []string{"College", "Social Sciences", "Beginner Psychology"}, 
		StudyBuddies:	pq.Int64Array {1},
	},
	models.Course{
		ID: 		157,
		Name: 		"Calculus 1",
		Keywords: 	[]string{"Derivatives", "Integrals", "Limits", "Continuous", "Velocity", "Chain Rule", "Logarithm", "Slope Field",},
		Categories: []string{"College", "Mathematics", "Calculus 1"}, 
		StudyBuddies:	pq.Int64Array {1,2,4},
	},
	models.Course{
		ID: 		140,
		Name: 		"Introduction to Computer Science",
		Keywords: 	[]string{"C++", "Functions", "Loops", "Primitive Types", "Return Statement", "Input/Output"},
		Categories: []string{"College", "Computer Science", "Introduction to Computer Science"}, 
		StudyBuddies:	pq.Int64Array {1,2,4},
	},
	models.Course{
		ID: 		141,
		Name: 		"Introduction to Bioengineering",
		Keywords: 	[]string{"Bioengineering"},
		Categories: []string{"College", "Bioengineering", "Introduction to Bioengineering"}, 
	},
	models.Course{
		ID: 		142,
		Name: 		"Communications I",
		Keywords: 	[]string{"Communications"},
		Categories: []string{"College", "Communications", "Communications I"}, 
	},
	models.Course{
		ID: 		143,
		Name: 		"Chemistry",
		Keywords: 	[]string{"Chemistry"},
		Categories: []string{"College", "Chemistry", "Organic Chemistry"}, 
	},
	models.Course{
		ID: 		144,
		Name: 		"Mechanical Physics",
		Keywords: 	[]string{"Physics"},
		Categories: []string{"College", "Physics", "Mechanical Physics"}, 
	},
	models.Course{
		ID: 		158,	
		Name: 		"AP Calculus AB",
		Keywords: 	[]string{"Derivatives", "Integrals", "Limits", "Continuous", "Velocity", "Chain Rule", "Logarithm", "Slope Field"},
		Categories: []string{"High School", "Mathematics", "AP Calculus AB"}, 
		StudyBuddies:	pq.Int64Array {7},
	},
}

var groups = []models.Group{
	models.Group{
		Name: 		"48ers",
		CourseName: "Human Anatomy",
		CourseID: 	155,
		AdminID: 	1,
		Members:	pq.Int64Array {1,2,5},
	},
	models.Group{
		Name: 		"LA People",
		CourseName: "Lower Division Linear Algebra",
		CourseID: 	115,
		AdminID: 	2,
		Members:	pq.Int64Array {1,2,4},
	},
	models.Group{
		Name: 		"LA Squad",
		CourseName: "Linear Algebra",
		CourseID: 	117,
		AdminID: 	5,
		Members:	pq.Int64Array {5,6},
	},
}

var listings = []models.Listing{
	models.Listing{
		ID:			99,
		CourseName: "Human Anatomy",
		Poster: 	1,
		CourseID: 	155,
		Tags:		[]string {"3.9+ GPA only", "No public school kids"},
		Description: "Looking for people to study with for upcoming final, need different study methods! Let's share! Movements are hard!",
	},
	models.Listing{
		ID:			100,
		CourseName: "Human Anatomy",
		Poster: 	2,
		CourseID: 	155,
		GroupName: "48ers",
		GroupID: 1,
		Tags:		[]string {"Casual"},
		Description: "Our group is looking for more people to study with for Human Anatomy! We are all West Coast Kids, come thru",
	},
	models.Listing{
		ID:			98,
		CourseName: "Linear Algebra",
		Poster: 	5,
		CourseID: 	117,
		GroupName: "LA Squad",
		GroupID: 3,
		Tags:		[]string {"Casual"},
		Description: "Our group is looking for more people to study with for Linear Algebra! We focus on the matrices",
	},
}

var invitations = []models.Invitation{
	models.Invitation{
		GroupName: 	"48ers",
		GroupID: 	1,
		ReceiveID: 	7,
		ReceiveName: "Thomas Brady",
		Type: 		true, //Request
		Status:		false,
	},
	models.Invitation{
		GroupName: 	"LA Squad",
		GroupID: 	3,
		ReceiveID: 	7,
		ReceiveName: "Thomas Brady",
		Type: 		true,
		Status:		false,
	},
}