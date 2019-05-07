package events

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mc303/gin-rest-api-sample/database/models"
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// type Rider = models.Rider

type JSON = common.JSON

// func list(c *gin.Context) {
// 	// Connection to the database
// 	db := c.MustGet("db").(*gorm.DB)
// 	// Close connection database
// 	defer db.Close()

// 	var event []models.Event
// 	// db.Set("gorm:auto_preload", true).Find(&event)
// 	// SELECT * FROM riders
// 	// db.Find(&event)vTypeOfRace
// 	db.Debug().Preload("Stages").Preload("TypeOfRaces").Find(&event)

// 	// Display JSON result
// 	c.JSON(200, event)
// }

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var events []models.Event

	// db.Debug().Preload("Stages").Preload("TypeOfRaces").Find(&event)
	// db.Model(&events).Related(&card)
	if cursor == "" {
		if err := db.Debug().Preload("Stages").Preload("TypeOfRaces").Find(&events).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Debug().Preload("Stages").Preload("TypeOfRaces").Where(condition, cursor).Find(&events).Error; err != nil {
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

// func read(c *gin.Context) {
// 	// Connection to the database
// 	db := c.MustGet("db").(*gorm.DB)
// 	// Close connection database
// 	defer db.Close()

// 	id := c.Params.ByName("id")

// 	var event models.Event
// 	// db.Set("gorm:auto_preload", true).Find(&event)
// 	// db.Debug().Preload("Stages").Find(&event, id)

// 	db.Debug().Preload("Stages").Preload("TypeOfRaces").Find(&event, id)

// 	if event.ID != 0 {
// 		// Display JSON result
// 		c.JSON(200, event)
// 	} else {
// 		// Display JSON error
// 		c.JSON(404, gin.H{"error": "Event not found"})
// 	}

// 	// Invoke-RestMethod -Uri "http://127.0.0.1:8080/api/v1/event/2" -Method get | ConvertTo-Json
// }

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var event models.Event

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Debug().Preload("Stages").Preload("TypeOfRaces").Find(&event, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, event.Serialize())
}
