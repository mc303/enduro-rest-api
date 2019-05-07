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
	RiderID     uint
	EventID    uint
	Events      Event `gorm:"ForeignKey:ID;association_foreignkey:EventsID"`
	ClassID     uint
	Classes     Class `gorm:"foreignkey:ID;association_foreignkey:ClassID"`
}

// Serialize serializes post data
func (p Registered) Serialize() common.JSON {
	return common.JSON{
		"ID":          p.ID,
		"Date":        p.Date,
		"StartNumber": p.StartNumber,
		"RidersID":    p.RidersID,
		"Riders":      p.Riders,
		"EventsID":    p.EventsID,
		"Events":      p.Events,
		"ClassID":     p.ClassID,
		"Class":       p.Class,
	}
}
