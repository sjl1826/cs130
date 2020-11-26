// invitation.go

package handlers

import (
	"cs130_back/models"
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	"time"
)

// InvitationByID gets the invitation by ID
func InvitationByID(db *gorm.DB, i *models.Invitation, w http.ResponseWriter) int {
	if err := i.GetInvitation(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Invitation not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return 0
	}
	return 1
}

// SendInvitationRequest required fields to create an invitation
type SendInvitationRequest struct {
	GroupName 		string 			`json:"group_name"`
	GroupID			int				`json:"group_id"`
	ReceiveID		int				`json:"receive_id"`
	ReceiveName 	string 			`json:"receive_name"`
	Type			bool			`json:"type"`
}

// SendInvitationResponse fields to send back
// HTTP status code 201 and group model in data
type SendInvitationResponse struct {
	ID          	int             `json:"g_id"`
	GroupName		string			`json:"group_name"`
	GroupID			int				`json:"group_id"`
	ReceiveID		int				`json:"receive_id"`
	ReceiveName 	string 			`json:"receive_name"`
	Type			bool			`json:"type"`
	Status			bool			`json:"status"`
	CreatedAt   	time.Time       `json:"CreatedAt"`
	UpdatedAt   	time.Time       `json:"UpdatedAt"`
}

func populateInvitationResponse(v *models.Invitation, r *CreateInvitationResponse) {
	r.ID = v.ID
	r.GroupName = v.GroupName
	r.GroupID = v.GroupID
	r.ReceiveID = v.ReceiveID
	r.ReceiveName = v.ReceiveName
	r.Type = v.Type
	r.Status = v.Status
	r.CreatedAt = v.CreatedAt
	r.UpdatedAt = v.UpdatedAt
}

// UpdateInvitationRequest for user invitation parsing
type UpdateInvitationRequest struct {
	ID				int    	`json:"u_id"`
	InvitationID	int 	`json:"invitation_id"`
	Status			string 	`json:"status"`
}

// UpdateInvitation will update the invitation for the specified user
func UpdateInvitation(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p UpdateInvitationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	invit := models.Invitation{ID: p.InvitationID, ReceiveID: p.ID}
	if InvitationByID(db, &invit, w) == 0 {
		return
	}
	
	if p.Status == "ACCEPT" {
		//add user to members of group
		group := models.Group{ID: invit.GroupID}
		if GetGroupByID(db, &group, w) == 0 {
			return
		}
		if err := group.AddMember(db, invit.ReceiveID); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} 

	//delete invitation after response
	if err := invit.DeleteInvitation(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}