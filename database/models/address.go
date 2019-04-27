package models

import (
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Address data model
type Address struct {
	// gorm.Model
	ID           uint
	Street       string
	StreetNumber string
	PostalCode   string
	City         string
	Country      string
	Riders       Rider `gorm:"ForeignKey:RidersID"`
	RidersID     uint
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
		"Riders":       p.Riders,
		"RidersID":     p.RidersID,
	}
}
