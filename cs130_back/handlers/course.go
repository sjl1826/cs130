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