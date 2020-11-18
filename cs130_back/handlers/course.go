// course.go

package handlers

import (
	"cs130_back/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// GetCourseByID gets the course by ID
func GetCourseByID(db *gorm.DB, c *models.Course, w http.ResponseWriter) int {
	if err := c.GetCourse(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Course not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return 0
	}
	return 1
}

// CreateCourseRequest required fields to create a course
type CreateCourseRequest struct {
	Description		string	`json:"description"`
	Title			string	`json:"title"`
}

// CreateCourseResponse fields to send back
// HTTP status code 201 and course model in data
type CreateCourseResponse struct {
	ID          	int 	`json:"id"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	Description		string	`json:"description"`
	Title			string	`json:"title"`
}

func populateCourseResponse(c *models.Course, r *CreateCourseResponse) {
	r.ID = c.ID
	r.CreatedAt = c.CreatedAt
	r.UpdatedAt = c.UpdatedAt
}

// CreateCourse initializes a new course in the database
func CreateCourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p CreateCourseRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	course := models.Course{}
	if err := course.CreateCourse(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if GetCourseByID(db, &course, w) == 0 {
		return
	}

	var cr CreateCourseResponse
	populateCourseResponse(&course, &cr)

	respondWithJSON(w, http.StatusCreated, cr)
}

// UpdateCourseRequest for course requests parsing
type UpdateCourseRequest struct {
	ID          	int 	`json:"id"`
	Description		string	`json:"description",omit_empty`
	Title			string	`json:"title",omit_empty`
}

// UpdateCourse will update the values of the specified course
func UpdateCourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p UpdateCourseRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	course := models.Course{ID: p.ID}
	if GetCourseByID(db, &course, w) == 0 {
		return
	}

	if err := course.UpdateCourse(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var cr CreateCourseResponse
	populateCourseResponse(&course, &cr)

	respondWithJSON(w, http.StatusOK, cr)
}

// DeleteCourse deletes the course permanently
func DeleteCourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := strconv.Atoi(vars["id"][0])
	if ok != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Course ID")
		return
	}

	course := models.Course{ID: id}
	if GetCourseByID(db, &course, w) == 0 {
		return
	}

	if err := course.DeleteCourse(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

type Subcategory struct {
	Name      string            `json:"name"`
	Keywords  []string          `json:"keywords"`
	Courses   []models.Course   `json:"courses"`
}

type Category struct {
	Name            string    		  `json:"category"`
	Courses         []models.Course   `json:"courses"`
	Subcategories   []Subcategory     `json:"subcategories"`
}

type CollegeResponse struct {
	Categories    []Category   `json:"categories"`
}

type HighSchoolResponse struct {
	Categories    []Category   `json:"categories"`
}

// CoursesResponse fields to send back
// HTTP status code 201 and user model in data
type CoursesResponse struct {
	College           CollegeResponse     `json:"College"`
	HighSchool        HighSchoolResponse  `json:"High School"`
}

func BuildCategoryResponse(db *gorm.DB, w http.ResponseWriter, r *http.Request, CG *[]Category, institution string) {
	var final []Category

	c := models.Course{}
	var categories []string
	if err := c.GetCategories(db, &categories, institution); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	for _, category := range categories {
		var courses []models.Course
		if err := c.GetCoursesByCategory(db, &courses, category, institution); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		cr := Category{Name: category, Courses: courses}

		// get the classes by subcategory for the category
		var subcategories []string
		if err := c.GetSubcategories(db, &subcategories, category, institution); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		for _, subcategory := range subcategories {
			var courses2 []models.Course
			var words []string
			if err := c.GetCoursesBySubcategory(db, &courses2, &words, subcategory, institution); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			subc := Subcategory{Name: subcategory, Courses: courses2, Keywords: words}
			cr.Subcategories = append(cr.Subcategories, subc)
		}

		final = append(final, cr)
	}
	(*CG) = final
}

type Test struct {
	Categories []string `json:"categories"`
}

// GetCourses retrieves and returns the master course list sorted by institution, categories, and subcategories
func GetCourses(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// var response CoursesResponse

	// BuildCategoryResponse(db, w, r, &response.College.Categories, "College")
	// BuildCategoryResponse(db, w, r, &response.HighSchool.Categories, "High School")

	var resp Test
	c := models.Course{}
	if err := c.GetCategories(db, &resp.Categories, "College"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}