package models

import (
	"time"

	"github.com/mc303/enduro-rest-api/lib/common"
)

// Run data model
type Run struct {
	ID        int `gorm:"primary_key"`
	StartTime time.Time
	EndTime   time.Time
	DNF       bool
	RiderID   int
	EventID   int
	StageID   int
}

// Riders    Rider `gorm:"foreignkey:RidersID;association_foreignkey:ID"`
// 	Stages    Stage `gorm:"ForeignKey:StagesID;association_foreignkey:ID"`
//	Events    Event `gorm:"ForeignKey:EventsID;association_foreignkey:ID"`

// Serialize serializes post data
func (p Run) Serialize() common.JSON {
	return common.JSON{
		"ID":        p.ID,
		"StartTime": p.StartTime,
		"EndTime":   p.EndTime,
		"DNF":       p.DNF,
		"RiderID":   p.RiderID,
		"EventID":   p.EventID,
		"StageID":   p.StageID,
	}
}

// "Stages":    p.Stages,
// "Events":    p.Events,
// "Riders":    p.Riders,
