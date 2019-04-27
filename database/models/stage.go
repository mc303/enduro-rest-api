package models

import "github.com/mc303/gin-rest-api-sample/lib/common"

// Stage data model
type Stage struct {
	ID       uint
	Name     string `gorm:"not null"`
	Location string `gorm:"not null"`
	Order    uint   `gorm:"not null"`
	Events   Event  `gorm:"ForeignKey:EventsID;association_foreignkey:ID"`
	EventsID uint
}

// Serialize serializes post data
func (p Stage) Serialize() common.JSON {
	return common.JSON{
		"ID":       p.ID,
		"Name":     p.Name,
		"Location": p.Location,
		"Oder":     p.Order,
		"Events":   p.Events,
		"EventsID": p.EventsID,
	}
}
