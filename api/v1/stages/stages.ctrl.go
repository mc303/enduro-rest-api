package stages

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

	var stages []models.Stage

	if cursor == "" {
		if err := db.Find(&stages).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Where(condition, cursor).Find(&stages).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var stage models.Stage

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Find(&stage, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, stage.Serialize())
}
