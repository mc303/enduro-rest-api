package stages

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

	var users []models.User

	if cursor == "" {
		if err := db.Find(&users).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Where(condition, cursor).Find(&users).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var user models.User

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Find(&user, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, user.Serialize())
}
