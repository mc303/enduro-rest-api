package models

import (
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Address data model
type Address struct {
	// gorm.Model
	ID           uint `gorm:"primary_key"`
	Street       string
	StreetNumber string
	PostalCode   string
	City         string
	Country      string
	PlusCode     string
	Coordinate   string
	RiderID      uint
	// EventsID     uint

	// Runs    []Run  `gorm:"ForeignKey:RiderID;association_foreignkey:ID"`
}

// Serialize serializes post data
func (p Address) Serialize() common.JSON {
	return common.JSON{
		"ID":           p.ID,
		"Street":       p.Street,
		"StreetNumber": p.StreetNumber,
		"PostalCode":   p.PostalCode,
		"City":         p.City,
		"Country":      p.Country,
		"PlusCode":     p.PlusCode,
		"Coordinate":   p.Coordinate,
		"RiderID":      p.RiderID,
		// "EventsID":     p.EventsID,
	}
}
