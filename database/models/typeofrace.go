package models

import "github.com/mc303/gin-rest-api-sample/lib/common"

// TypeOfRace data model
type TypeOfRace struct {
	ID   uint   `gorm:"primary_key"` //	gorm.Model
	Name string `gorm:"not null"`
	// EventsID uint
}

// Serialize serializes post data
func (p TypeOfRace) Serialize() common.JSON {
	return common.JSON{
		"ID":   p.ID,
		"Name": p.Name,
		// "EventsID": p.EventsID,
	}
}
