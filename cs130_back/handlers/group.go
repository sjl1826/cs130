// group.go

package handlers

import (
	"cs130_back/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

// GetGroupByID gets the group by ID
func GetGroupByID(db *gorm.DB, g *models.Group, w http.ResponseWriter) int {
	if err := g.GetGroup(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Group not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return 0
	}
	return 1
}

// CreateGroupRequest required fields to create a group
type CreateGroupRequest struct {
	AdminID  int    `json:"admin_id"`
	Name     string `json:"name"`
	CourseID int    `json:"course_id"`
}

// CreateResponse fields to send back
// HTTP status code 201 and group model in data
type CreateGroupResponse struct {
	ID        int       `json:"u_id"`
	AdminID   int       `json:"admin_id"`
	Name      string    `json:"name"`
	CourseID  int       `json:"course_id"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func populateGroupResponse(g *models.Group, r *CreateGroupResponse) {
	r.ID = g.ID
	r.AdminID = g.AdminID
	r.Name = g.Name
	r.CourseID = g.CourseID
	r.CreatedAt = g.CreatedAt
	r.UpdatedAt = g.UpdatedAt
}

// CreateGroup initializes a new group in the database
func CreateGroup(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p CreateGroupRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	group := models.Group{Name: p.Name, CourseID: p.CourseID, AdminID: p.AdminID}
	if err := group.CreateGroup(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := group.GetGroup(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Group not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	var gr CreateGroupResponse
	populateGroupResponse(&group, &gr)

	respondWithJSON(w, http.StatusCreated, gr)
}
