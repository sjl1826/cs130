// subject.go

package handlers

import (
	"cs130_back/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// GetSubjectByID gets the subject by ID
func GetSubjectByID(db *gorm.DB, s *models.Subject, w http.ResponseWriter) int {
	if err := s.GetSubject(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Subject not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return 0
	}
	return 1
}

// CreateSubjectRequest required fields to create a subject
type CreateSubjectRequest struct {
	Description		string	`json:"description"`
	Title			string	`json:"title"`
}

// CreateSubjectResponse fields to send back
// HTTP status code 201 and subject model in data
type CreateSubjectResponse struct {
	ID          	int 	`json:"id"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	Description		string	`json:"description"`
	Title			string	`json:"title"`
}

func populateSubjectResponse(s *models.Subject, r *CreateSubjectResponse) {
	r.ID = s.ID
	r.CreatedAt = s.CreatedAt
	r.UpdatedAt = s.UpdatedAt
	r.Description = s.Description
	r.Title = s.Title
}

// CreateSubject initializes a new subject in the database
func CreateSubject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p CreateSubjectRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	subject := models.Subject{Description: p.Description, Title: p.Title}
	if err := subject.CreateSubject(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if GetSubjectByID(db, &subject, w) == 0 {
		return
	}

	var cr CreateSubjectResponse
	populateSubjectResponse(&subject, &cr)

	respondWithJSON(w, http.StatusCreated, cr)
}

// UpdateSubjectRequest for subject requests parsing
type UpdateSubjectRequest struct {
	ID          	int 	`json:"id"`
	Description		string	`json:"description",omit_empty`
	Title			string	`json:"title",omit_empty`
}

// UpdateSubject will update the values of the specified subject
func UpdateSubject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p UpdateSubjectRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	subject := models.Subject{ID: p.ID}
	if GetSubjectByID(db, &subject, w) == 0 {
		return
	}

	if p.Description != "" {
		subject.Description = p.Description
	}
	if p.Title != "" {
		subject.Title = p.Title
	}

	if err := subject.UpdateSubject(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var cr CreateSubjectResponse
	populateSubjectResponse(&subject, &cr)

	respondWithJSON(w, http.StatusOK, cr)
}

// DeleteSubject deletes the subject permanently
func DeleteSubject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := strconv.Atoi(vars["id"][0])
	if ok != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Subject ID")
		return
	}

	subject := models.Subject{ID: id}
	if GetSubjectByID(db, &subject, w) == 0 {
		return
	}

	if err := subject.DeleteSubject(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}