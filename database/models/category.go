package models

import (
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Category data model
type Category struct {
	//gorm.Model
	ID     int
	Name   string `gorm:"not null"`
	Gender int
}

// Serialize serializes post data
func (p Category) Serialize() common.JSON {
	return common.JSON{
		"ID":     p.ID,
		"Name":   p.Name,
		"Gender": p.Gender,
	}
}
