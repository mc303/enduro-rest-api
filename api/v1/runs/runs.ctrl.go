package runs

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

	var run []models.Run
	// db.Set("gorm:auto_preload", true).Find(&event)
	// SELECT * FROM riders
	// db.Find(&event)
	db.Debug().Preload("Riders.Class").Preload("Stages.Events").Find(&run)

	// Display JSON result
	c.JSON(200, run)
}

// func read(c *gin.Context) {
// 	// Connection to the database
// 	db := c.MustGet("db").(*gorm.DB)
// 	// Close connection database
// 	// defer db.Close()

// 	id := c.Params.ByName("id")

// 	var run models.Run

// 	db.Debug().Preload("Riders.Class").Preload("Stages.Events").Find(&run, id)

// 	if run.ID != 0 {
// 		// Display JSON result
// 		c.JSON(200, run)
// 	} else {
// 		// Display JSON error
// 		c.JSON(404, gin.H{"error": "Run not found"})
// 	}

// 	//
// }

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var run models.Run

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Debug().Preload("Riders.Class").Preload("Stages.Events").Find(&run, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, run.Serialize())
}

func readstage(c *gin.Context) {
	// Connection to the database
	db := c.MustGet("db").(*gorm.DB)
	// Close connection database
	// defer db.Close()

	id := c.Params.ByName("id")

	var run []models.Run
	//db.Set("gorm:auto_preload", true).Find(&stage)
	db.Debug().Preload("Riders.Class").
		Preload("Stages.Events").
		Where("stages_id = ?", id).
		Find(&run)

	// Display JSON result
	c.JSON(200, run)

	//
}

func readevent(c *gin.Context) {
	// Connection to the database
	db := c.MustGet("db").(*gorm.DB)
	// Close connection database
	// defer db.Close()

	id := c.Params.ByName("id")

	//var run models.Run

	var run []models.Run

	db.Debug().Preload("Riders.Class").Joins("JOIN stages on runs.stages_id = stages.id").
		Where("stages.events_id = ?", id).
		Preload("Stages.Events").Find(&run)

	// Display JSON result
	c.JSON(200, run)

	// Invoke-RestMethod -Uri "http://127.0.0.1:8080/api/v1/runs/event/1" -Method get | ConvertTo-Json
}

func readrider(c *gin.Context) {
	// Connection to the database
	db := c.MustGet("db").(*gorm.DB)
	// Close connection database
	// defer db.Close()

	id := c.Params.ByName("id")

	var run []models.Run
	//db.Set("gorm:auto_preload", true).Find(&stage)
	db.Debug().Preload("Riders.Class").
		Preload("Stages.Events").
		Where("riders_id = ?", id).
		Find(&run)

	// Display JSON result
	c.JSON(200, run)

	// Invoke-RestMethod -Uri "http://127.0.0.1:8080/api/v1/runs/rider/1" -Method get | ConvertTo-Json
}
