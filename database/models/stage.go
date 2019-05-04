package models

import "github.com/mc303/gin-rest-api-sample/lib/common"

// Stage data model
type Stage struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"not null"`
	Location string `gorm:"not null"`
	Order    uint   `gorm:"not null"`
	EventsID uint
	// Events   Event  `gorm:"ForeignKey:EventsID;association_foreignkey:ID"`

}

// Serialize serializes post data
func (p Stage) Serialize() common.JSON {
	return common.JSON{
		"ID":       p.ID,
		"Name":     p.Name,
		"Location": p.Location,
		"Oder":     p.Order,
		"EventsID": p.EventsID,
	}
}
