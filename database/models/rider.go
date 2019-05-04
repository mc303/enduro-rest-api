package models

import (
	"time"

	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Rider data model
type Rider struct {
	// gorm.Model
	ID          uint      `gorm:"primary_key"`
	Firstname   string    `gorm:"not null"`
	Lastname    string    `gorm:"not null"`
	Gender      int       `gorm:"not null"`
	Mail        string    `gorm:"not null"`
	Birthday    time.Time `gorm:"not null"`
	AddressesID uint
	Addresses   Address `gorm:"ForeignKey:ID;association_foreignkey:AddressesID"`
	RunsID      uint
	Runs        []Run `gorm:"ForeignKey:RidersID;association_foreignkey:RunsID"`
	ResultsID   uint
	Results     Result `gorm:"ForeignKey:RidersID;association_foreignkey:ResultsID"`
	StagesID    uint
	Stages      []Stage `gorm:"ForeignKey:ID;association_foreignkey:StagesID"`
}

// Serialize serializes post data
func (p Rider) Serialize() common.JSON {
	return common.JSON{
		"ID":          p.ID,
		"Firstname":   p.Firstname,
		"Lastname":    p.Lastname,
		"Gender":      p.Gender,
		"Mail":        p.Mail,
		"Birthday":    p.Birthday,
		"AddressesID": p.AddressesID,
		"Addresses":   p.Addresses,
		"Runs":        p.Runs,
		"RunsID":      p.RunsID,
		"ResultsID":   p.ResultsID,
		"Results":     p.Results,
		"StagesID":    p.StagesID,
		"Stages":      p.Stages,
	}
}
