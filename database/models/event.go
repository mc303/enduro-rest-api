package models

import (
	"time"

	"github.com/mc303/enduro-rest-api/lib/common"
)

// Event data model
type Event struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Date      time.Time `gorm:"not null"`
	Addresses Address
	Races     []Race

	// Stages        []Stage `gorm:"ForeignKey:EventsID"`
}

// Serialize serializes post data
func (p Event) Serialize() common.JSON {
	return common.JSON{
		"ID":        p.ID,
		"Name":      p.Name,
		"Date":      p.Date,
		"Addresses": p.Addresses,
		"Races":     p.Races,
	}
}
