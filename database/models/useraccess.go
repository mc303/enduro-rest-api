package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mc303/enduro-rest-api/lib/common"
)

// UserAccess data model
type UserAccess struct {
	gorm.Model
	SourceID int
	Role     string `gorm:"not null"`
	// ID           int   `gorm:"primary_key"`
	// Access       int `gorm:"not null"`
}

// Serialize serializes post data
func (p UserAccess) Serialize() common.JSON {
	return common.JSON{
		"ID":         p.ID,
		"SourceID":   p.SourceID,
		"Role":       p.Role,
		"created_at": p.CreatedAt,
		// "Access":       p.Access,
	}
}
