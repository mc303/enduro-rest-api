package models

import "github.com/mc303/gin-rest-api-sample/lib/common"

// Stage data model
type Stage struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Location  string `gorm:"not null"`
	Order     int    `gorm:"not null"`
	EventID   int
	Addresses Address
	Runs      []Run
	RaceID    int

	// Events   Event  `gorm:"ForeignKey:EventsID;association_foreignkey:ID"`

}

// Serialize serializes post data
func (p Stage) Serialize() common.JSON {
	return common.JSON{
		"ID":        p.ID,
		"Name":      p.Name,
		"Location":  p.Location,
		"Order":     p.Order,
		"Addresses": p.Addresses,
		"EventID":   p.EventID,
		"Runs":      p.Runs,
	}
}
