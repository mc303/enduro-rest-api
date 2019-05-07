package models

import (
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Event data model
type Event struct {
	ID             uint   `gorm:"primary_key"`
	Name           string `gorm:"not null"`
	Date           string `gorm:"not null"`
	TotalStages    int    `gorm:"not null"`
	AddressesID    uint
	Addresses      Address `gorm:"foreignkey:ID;association_foreignkey:AddressesID"`
	TypeOfRacesID  uint
	TypeOfRaces    TypeOfRace `gorm:"foreignkey:ID;association_foreignkey:TypeOfRacesID"` // One-To-One relationship (has one)
	SeasonsID      uint
	Seasons        Season `gorm:"foreignkey:ID;association_foreignkey:SeasonsID"` //;association_foreignkey:ID"
	Stages         []Stage // `gorm:"foreignkey:EventsID;association_foreignkey:StagesEventsID"` // ;association_foreignkey:ID
	Results        []Result
	Registereds    []Registered
	// Stages        []Stage `gorm:"ForeignKey:EventsID"`
}

// Serialize serializes post data
func (p Event) Serialize() common.JSON {
	return common.JSON{
		"ID":             p.ID,
		"Name":           p.Name,
		"Date":           p.Date,
		"TotalStages":    p.TotalStages,
		"Addresses":      p.AddressesID,
		"AddressesID":    p.AddressesID,
		"TypeOfRaces":    p.TypeOfRaces,
		"TypeOfRacesID":  p.TypeOfRacesID,
		"Seasons":        p.Seasons,
		"SeasonsID":      p.SeasonsID,
		"StagesEventsID": p.StagesEventsID,
		"Stages":         p.Stages,
	}
}
