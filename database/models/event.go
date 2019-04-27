package models

import (
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Event data model
type Event struct {
	ID            uint
	Name          string `gorm:"not null"`
	AddressID     uint
	Date          string     `gorm:"not null"`
	TotalStages   int        `gorm:"not null"`
	TypeOfRaces   TypeOfRace `gorm:"foreignkey:TypeOfRacesID;association_foreignkey:ID"` // One-To-One relationship (has one)
	TypeOfRacesID uint       //`gorm:"ForeignKey:id"`
	Stages        []Stage    `gorm:"ForeignKey:EventsID"`
}

// Serialize serializes post data
func (p Event) Serialize() common.JSON {
	return common.JSON{
		"ID":            p.ID,
		"Name":          p.Name,
		"AddressID":     p.AddressID,
		"Date":          p.Date,
		"TotalStages":   p.TotalStages,
		"TypeOfRaces":   p.TypeOfRaces,
		"TypeOfRacesID": p.TypeOfRacesID,
		"Stages":        p.Stages,
	}
}
