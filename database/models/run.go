package models

import (
	"time"

	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Run data model
type Run struct {
	ID        uint
	StartTime time.Time
	EndTime   time.Time
	DNF       bool
	Riders    Rider `gorm:"foreignkey:RidersID;association_foreignkey:ID"`
	RidersID  uint  // `gorm:"ForeignKey:id"`
	Stages    Stage `gorm:"ForeignKey:StagesID;association_foreignkey:ID"`
	StagesID  uint
}

// Serialize serializes post data
func (p Run) Serialize() common.JSON {
	return common.JSON{
		"ID":        p.ID,
		"StartTime": p.StartTime,
		"EndTime":   p.EndTime,
		"DNF":       p.DNF,
		"Riders":    p.Riders,
		"RidersID":  p.RidersID,
		"Stages":    p.Stages,
		"StagesID":  p.StagesID,
	}
}
