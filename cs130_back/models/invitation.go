// invitation.go

package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// Invitation structure
type Invitation struct {
	ID          	int 			`gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	GroupName		string			`json:"group_name"`
	GroupID			int				`json:"group_id"`
	ReceiveID		int				`json:"receive_id"`
	Type			bool			`json:"type"`
	Status			bool			`json:"status"`
}

// CreateInvitation creates a new invitation object in database
func (r *Invitation) CreateInvitation(db *gorm.DB) error {
	now := time.Now()
	r.CreatedAt = now
	retVal := db.Create(r).Table("invitations").Scan(&r)
	return retVal.Error
}

// UpdateInvitation updates invitation fields
func (r *Invitation) UpdateInvitation(db *gorm.DB) error {
	now := time.Now()
	r.UpdatedAt = now
	retVal := db.Save(&r).Table("invitations")
	return retVal.Error
}

// GetInvitation retrieves invitation object from database
func (r *Invitation) GetInvitation(db *gorm.DB) error {
	retVal := db.First(&r, r.ID).Scan(&r).Table("invitations")
	return retVal.Error
}

// DeleteInvitation deletes invitation from database
func (r *Invitation) DeleteInvitation(db *gorm.DB) error {
	retVal := db.Exec("DELETE FROM invitations WHERE ID=" + strconv.Itoa(r.ID))
	return retVal.Error
}
