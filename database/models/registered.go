package models

import (
	"time"

	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Registered data model
type Registered struct {
	ID          uint `gorm:"primary_key"`
	Date        time.Time
	StartNumber string
	RidersID    uint
	Riders      Rider `gorm:"ForeignKey:ID;association_foreignkey:RidersID"`
	EventsID    uint
	Events      Event `gorm:"ForeignKey:ID;association_foreignkey:EventsID"`
	ClassID     uint
	Class       Class `gorm:"foreignkey:ID;association_foreignkey:ClassID"`
}

// Serialize serializes post data
func (p Registered) Serialize() common.JSON {
	return common.JSON{
		"ID":          p.ID,
		"Date":        p.Date,
		"Riders":      p.Riders,
		"RidersID":    p.RidersID,
		"Events":      p.Events,
		"EventsID":    p.EventsID,
		"StartNumber": p.StartNumber,
	}
}
