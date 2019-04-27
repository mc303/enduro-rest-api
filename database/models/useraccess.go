package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// UserAccess data model
type UserAccess struct {
	gorm.Model
	SourceID uint
	Role     string `gorm:"not null"`
	// ID           uint   `gorm:"primary_key"`
	// Access       uint `gorm:"not null"`
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
