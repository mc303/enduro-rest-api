package models

import (
	"time"

	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Race data model
type Race struct {
	ID            int       `gorm:"primary_key"`
	Name          string    `gorm:"not null"`
	Date          time.Time `gorm:"not null"`
	TotalStages   int       `gorm:"not null"`
	DisciplinesID int
	Disciplines   Discipline `gorm:"foreignkey:ID;association_foreignkey:DisciplinesID"` // One-To-One relationship (has one)
	Stages        []Stage    // `gorm:"foreignkey:EventsID;association_foreignkey:StagesEventsID"` // ;association_foreignkey:ID
	Runs          []Run
	Results       []Result
	Registereds   []Registered
	EventID       int
	// Stages        []Stage `gorm:"ForeignKey:EventsID"`
}

// Serialize serializes post data
func (p Race) Serialize() common.JSON {
	return common.JSON{
		"ID":            p.ID,
		"Name":          p.Name,
		"Date":          p.Date,
		"TotalStages":   p.TotalStages,
		"DisciplinesID": p.DisciplinesID,
		"Disciplines":   p.Disciplines,
		"Stages":        p.Stages,
		"Results":       p.Results,
		"Registereds":   p.Registereds,
	}
}
