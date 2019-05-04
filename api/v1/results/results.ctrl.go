package results

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mc303/gin-rest-api-sample/database/models"
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// type Result struct {
// 	ID        uint
// 	TotalTime time.Time
// 	Riders    Rider `gorm:"ForeignKey:RidersID;association_foreignkey:ID"`
// 	RidersID  uint
// 	Events    Event `gorm:"ForeignKey:EventsID;association_foreignkey:ID"`
// 	EventsID  uint
// 	Seasons   Season `gorm:"ForeignKey:SeasonsID;association_foreignkey:ID"`
// 	SeasonsID uint
// 	Place     uint
// }

// type Rider = models.Rider

type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var results []models.Result

	// db.Debug().Preload("Stages").Preload("TypeOfRaces").Find(&event)

	if cursor == "" {
		if err := db.Preload("Events").Preload("Events.TypeOfRaces").Preload("Events.Stages").Preload("Riders").Preload("Seasons").Find(&results).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Preload("Events").Preload("Events.TypeOfRaces").Preload("Events.Stages").Preload("Riders").Preload("Seasons").Where(condition, cursor).Find(&results).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(results)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = results[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var results models.Result

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Preload("Events").Preload("Riders").Find(&results, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, results.Serialize())
}
