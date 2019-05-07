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
	Addresses   Address
	Runs        []Run
	Results     []Result
	Stages      []Stage
	Registereds []Registered
}

// Serialize serializes post data
func (p Rider) Serialize() common.JSON {
	return common.JSON{
		"ID":        p.ID,
		"Firstname": p.Firstname,
		"Lastname":  p.Lastname,
		"Gender":    p.Gender,
		"Mail":      p.Mail,
		"Birthday":  p.Birthday,
		"Addresses": p.Addresses,
		"Runs":      p.Runs,
		"Results":   p.Results,
		"Stages":    p.Stages,
	}
}
