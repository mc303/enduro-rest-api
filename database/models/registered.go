package models

import (
	"time"

	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Registered data model
type Registered struct {
	ID          uint
	Date        time.Time
	Riders      Rider `gorm:"ForeignKey:RidersID;association_foreignkey:ID"`
	RidersID    uint
	Events      Event `gorm:"ForeignKey:EventsID;association_foreignkey:ID"`
	EventsID    uint
	StartNumber string
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
