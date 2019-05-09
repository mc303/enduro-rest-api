package events

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mc303/enduro-rest-api/database/models"
	"github.com/mc303/enduro-rest-api/lib/common"
)

// type Rider = models.Rider

type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var events []models.Event

	// db.Preload("Stages").Preload("TypeOfRaces").Find(&event)
	// db.Model(&events).Related(&card)
	if cursor == "" {
		if err := db.Find(&events).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		// db.Preload("Stages").Preload("TypeOfRaces").Where(condition, cursor).Find(&events)
		if err := db.Where(condition, cursor).Find(&events).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(events)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = events[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var event models.Event

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Find(&event, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, event.Serialize())
}
