package seasons

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mc303/gin-rest-api-sample/database/models"
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// type Rider = models.Rider

type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var seasons []models.Season

	// db.Debug().Preload("Stages").Preload("TypeOfRaces").Find(&event)

	if cursor == "" {
		if err := db.Debug().Find(&seasons).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Debug().Where(condition, cursor).Find(&seasons).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(seasons)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = seasons[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var seasons models.Season

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Debug().Find(&seasons, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, seasons.Serialize())
}
