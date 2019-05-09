package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mc303/enduro-rest-api/lib/common"
)

// Post data model
type Post struct {
	gorm.Model
	Text   string `sql:"type:text;"`
	User   User   `gorm:"foreignkey:UserID"`
	UserID uint
}

// Serialize serializes post data
func (p Post) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"text":       p.Text,
		"user":       p.User.Serialize(),
		"created_at": p.CreatedAt,
	}
}
