package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// User data model
type User struct {
	gorm.Model
	Username     string `gorm:"column:username;not null"`
	PasswordHash string `gorm:"column:password;not null"`
	DisplayName  string `gorm:"not null"`
	UUID         string
	// ID           int   `gorm:"primary_key"`
	// Access       int `gorm:"not null"`
}

// Serialize serializes post data
func (p User) Serialize() common.JSON {
	return common.JSON{
		"ID":           p.ID,
		"Username":     p.Username,
		"PasswordHash": p.PasswordHash,
		"DisplayName":  p.DisplayName,
		"UUID":         p.UUID,
		"created_at":   p.CreatedAt,
		// "Access":       p.Access,
	}
}

// Read serializes post data
func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
	u.PasswordHash = m["password"].(string)
	u.DisplayName = m["display_name"].(string)
	u.UUID = m["uuid"].(string)
	u.CreatedAt = m["created_at"].(time.Time)
}
