package models

import "github.com/mc303/enduro-rest-api/lib/common"

// Discipline data model
type Discipline struct {
	ID   int    `gorm:"primary_key"` //	gorm.Model
	Name string `gorm:"not null"`
	// EventsID int
}

// Serialize serializes post data
func (p Discipline) Serialize() common.JSON {
	return common.JSON{
		"ID":   p.ID,
		"Name": p.Name,
		// "EventsID": p.EventsID,
	}
}
