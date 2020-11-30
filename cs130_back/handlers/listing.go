package handlers

import (	
	"cs130_back/models"
	"encoding/json"
	"net/http"
	"time"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// GetListingByID gets the listing by ID
func GetListingByID(db *gorm.DB, l *models.Listing, w http.ResponseWriter) int {
	if err := l.GetListing(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Listing not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return 0
	}
	return 1
}

//CreateListingRequest provides fields required to create a listing
type CreateListingRequest struct {
	Poster		int 			`json:"poster"`
	CourseID  	int 			`json:"course_id"`
	Description	string 			`json:"text_description"`
	Tags		pq.StringArray 	`json:"tags"`
}

//CreateListingResponse provides fields sent back
type CreateListingResponse struct {
	ID          	int 			`json:"id"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	CourseName		string			`json:"course_name"`
	Poster			int				`json:"poster"`
	CourseID		int				`json:"course_id"`
	Description		string			`json:"text_description"`
	GroupID			int				`json:"group_id"`		//optional group		
	Tags			pq.StringArray 	`json:"tags"`
}

func populateListingResponse(l *models.Listing, r *CreateListingResponse) {
	r.ID = l.ID
	r.CreatedAt = l.CreatedAt
	r.UpdatedAt = l.UpdatedAt
	r.CourseName = l.CourseName
	r.Poster = l.Poster
	r.CourseID = l.CourseID
	r.Description = l.Description
	r.GroupID = l.GroupID
	r.Tags = l.Tags
}

//CreateListing creates a new Listing in the database, and adds it to user and course
func CreateListing(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p CreateListingRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	listing := models.Listing{
		Poster: p.Poster, CourseID: p.CourseID, Description: p.Description, Tags: p.Tags}

	if err := listing.CreateListing(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if GetListingByID(db, &listing, w) == 0 {
		return
	}

	var clr CreateListingResponse
	populateListingResponse(&listing, &clr)

	respondWithJSON(w, http.StatusCreated, clr)
}

//GetListing retrieves a listing
func GetListing(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := strconv.Atoi(vars["id"][0])
	if ok != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid listing id")
		return
	}

	listing := models.Listing{ID: id}
	if GetListingByID(db, &listing, w) == 0 {
		return
	}

	var clr CreateListingResponse	
	populateListingResponse(&listing, &clr)

	respondWithJSON(w, http.StatusCreated, clr)
}

//UpdateListingRequest holds listing fields that can be updated
type UpdateListingRequest struct {
	ID          	int 			`json:"id"`
	Description		string			`json:"text_description"`
	GroupID			int				`json:"group_id"`		//optional group (set to 0 if not used)
	Tags			pq.StringArray 	`json:"tags"`
}


//UpdateListing updates a listing
func UpdateListing(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p UpdateListingRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	listing := models.Listing{ID: p.ID}
	if GetListingByID(db, &listing, w) == 0 {
		return
	}

	//Only update fields if they are non empty
	if p.Description != "" {
		listing.Description = p.Description
	}

	if p.GroupID != 0 {
		listing.GroupID = p.GroupID
	}

	if p.Tags != nil {
		listing.Tags = p.Tags
	}

	if err := listing.UpdateListing(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var clr CreateListingResponse
	populateListingResponse(&listing, &clr)

	respondWithJSON(w, http.StatusOK, clr)
}

//DeleteListing deletes a listing and removes the listing from user and course
func DeleteListing(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := strconv.Atoi(vars["id"][0])
	if ok != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid listing id")
		return
	}

	listing := models.Listing{ID: id}
	if GetListingByID(db, &listing, w) == 0 {
		return
	}

	//Delete listing object
	if err := listing.DeleteListing(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}