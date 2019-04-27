package models

import (
	"time"

	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Rider data model
type Rider struct {
	// gorm.Model
	ID        uint
	Firstname string    `gorm:"not null"`
	Lastname  string    `gorm:"not null"`
	Gender    int       `gorm:"not null"`
	Mail      string    `gorm:"not null"`
	Class     Class     `gorm:"foreignkey:ClassID;association_foreignkey:ID"` // One-To-One relationship (has one)
	ClassID   uint      //`gorm:"ForeignKey:id"`
	Birthday  time.Time `gorm:"not null"`
	// Addresses []Address `gorm:"foreignkey:RidersID"`
	// Runs    []Run  `gorm:"ForeignKey:RiderID"`
}

// Serialize serializes post data
func (p Rider) Serialize() common.JSON {
	return common.JSON{
		"ID":        p.ID,
		"Firstname": p.Firstname,
		"Lastname":  p.Lastname,
		"Gender":    p.Gender,
		"Mail":      p.Mail,
		"Class":     p.Class,
		"ClassID":   p.ClassID,
		"Birthday":  p.Birthday,
		// "Addresses": p.Addresses,
	}
}
