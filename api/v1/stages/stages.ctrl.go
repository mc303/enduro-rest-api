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
	// Connection to the database
	db := c.MustGet("db").(*gorm.DB)
	// Close connection database
	// defer db.Close()

	var stage []models.Stage
	// db.Set("gorm:auto_preload", true).Find(&event)
	// SELECT * FROM riders
	// db.Find(&event)
	db.Debug().Preload("Events").Find(&stage)

	// Display JSON result
	c.JSON(200, stage)
}

func read(c *gin.Context) {
	// Connection to the database
	db := c.MustGet("db").(*gorm.DB)
	// Close connection database
	// defer db.Close()

	id := c.Params.ByName("id")

	var stage models.Stage
	//db.Set("gorm:auto_preload", true).Find(&stage)
	db.Debug().Preload("Events").Find(&stage, id)

	if stage.ID != 0 {
		// Display JSON result
		c.JSON(200, stage)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Stage not found"})
	}
	// Invoke-RestMethod -Uri "http://127.0.0.1:8080/api/v1/stage/3" -Method get | ConvertTo-Json
}
