package models

import (
	"time"

	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Result data model
type Result struct {
	ID        uint
	TotalTime time.Time
	Riders    Rider `gorm:"ForeignKey:RidersID;association_foreignkey:ID"`
	RidersID  uint
	Events    Event `gorm:"ForeignKey:EventsID;association_foreignkey:ID"`
	EventsID  uint
	Place     uint
}

// Serialize serializes post data
func (p Result) Serialize() common.JSON {
	return common.JSON{
		"ID":        p.ID,
		"TotalTime": p.TotalTime,
		"Riders":    p.Riders,
		"RidersID":  p.RidersID,
		"Events":    p.Events,
		"EventsID":  p.EventsID,
		"Place":     p.Place,
	}
}
