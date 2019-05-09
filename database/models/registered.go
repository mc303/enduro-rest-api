package models

import (
	"time"

	"github.com/mc303/enduro-rest-api/lib/common"
)

// Registered data model
type Registered struct {
	ID           int `gorm:"primary_key"`
	Date         time.Time
	StartNumber  string
	RiderID      int
	EventID      int
	CategoriesID int
	Categories   Category `gorm:"foreignkey:ID;association_foreignkey:CategoriesID"`
	// Events      Event `gorm:"ForeignKey:ID;association_foreignkey:EventsID"`

}

// Serialize serializes post data
func (p Registered) Serialize() common.JSON {
	return common.JSON{
		"ID":           p.ID,
		"Date":         p.Date,
		"StartNumber":  p.StartNumber,
		"RiderID":      p.RiderID,
		"EventID":      p.EventID,
		"CategoriesID": p.CategoriesID,
		"Categories":   p.Categories,
	}
}
