package models

import (
	"time"

	"github.com/mc303/enduro-rest-api/lib/common"
)

// Result data model
type Result struct {
	ID        int `gorm:"primary_key"`
	TotalTime time.Time
	Place     int
	RiderID   int
	EventID   int
	StageID   int
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
