package disciplines

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mc303/gin-rest-api-sample/database/models"
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// type Rider = models.Rider

// JSON some text
type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var disciplines []models.Discipline

	// db.Preload("Stages").Preload("TypeOfRaces").Find(&event)

	if cursor == "" {
		if err := db.Find(&disciplines).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Where(condition, cursor).Find(&disciplines).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(disciplines)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = disciplines[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var discipline models.Discipline

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Find(&discipline, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, discipline.Serialize())
}
