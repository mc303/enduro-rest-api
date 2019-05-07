package models

import (
	"time"

	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Run data model
type Run struct {
	ID        uint `gorm:"primary_key"`
	StartTime time.Time
	EndTime   time.Time
	DNF       bool
	RiderID   uint
	EventsID  uint
	StagesID  uint
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
		"RidersID":  p.RidersID,
		"EventsID":  p.EventsID,
		"StagesID":  p.StagesID,
	}
}

// "Stages":    p.Stages,
// "Events":    p.Events,
// "Riders":    p.Riders,
