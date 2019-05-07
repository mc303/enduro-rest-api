package models

import "github.com/mc303/gin-rest-api-sample/lib/common"

// Season data model
type Season struct {
	ID   uint   `gorm:"primary_key"` //	gorm.Model
	Name string `gorm:"not null"`
	Year uint   `gorm:"not null"`
	// EventsID uint
}

// Serialize serializes post data
func (p Season) Serialize() common.JSON {
	return common.JSON{
		"ID":   p.ID,
		"Name": p.Name,
		"Year": p.Year,
		// "EventsID": p.EventsID,
	}
}
