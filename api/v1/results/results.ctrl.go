package results

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mc303/gin-rest-api-sample/database/models"
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var results []models.Result

	// db.Preload("Stages").Preload("TypeOfRaces").Find(&event)
	//  db.Preload("Events").Preload("Events.TypeOfRaces").Preload("Events.Stages").Preload("Riders").Preload("Seasons")
	if cursor == "" {
		if err := db.Find(&results).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Where(condition, cursor).Find(&results).Error; err != nil {
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
	if err := db.Find(&results, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, results.Serialize())
}
