package models

import (
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// Class data model
type Class struct {
	//gorm.Model
	ID   uint
	Name string `gorm:"not null"`
}

// Serialize serializes post data
func (p Class) Serialize() common.JSON {
	return common.JSON{
		"ID":   p.ID,
		"Name": p.Name,
	}
}
