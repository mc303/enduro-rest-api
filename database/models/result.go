package models

import (
	"time"

	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Result data model
type Result struct {
	ID        uint `gorm:"primary_key"`
	TotalTime time.Time
	Place     uint
	RiderID   uint
	EventID   uint
	// Events    Event `gorm:"ForeignKey:ID;association_foreignkey:EventsID"`
}

// Serialize serializes post data
func (p Result) Serialize() common.JSON {
	return common.JSON{
		"ID":        p.ID,
		"TotalTime": p.TotalTime,
		"Place":     p.Place,
		"RiderID":   p.RiderID,
		"EventID":   p.EventID,
		// "Events":    p.Events,
	}
}
