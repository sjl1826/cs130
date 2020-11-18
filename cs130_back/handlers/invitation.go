// invitation.go

package handlers

import (
	"cs130_back/models"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

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
		invit.Status = true
	} else {
		invit.Status = false
	}

	if err := invit.UpdateInvitation(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}