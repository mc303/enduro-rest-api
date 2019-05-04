package riders

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mc303/gin-rest-api-sample/database/models"
	"github.com/mc303/gin-rest-api-sample/lib/common"
)

// type Rider = models.Rider

type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//db := InitDb()
	// defer db.Close()

	var rider models.Rider
	c.Bind(&rider)

	if rider.Firstname != "" && rider.Lastname != "" {
		// INSERT INTO "users" (name) VALUES (user.Name);
		db.Create(&rider)
		// Display error
		c.JSON(201, gin.H{"success": rider})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func list(c *gin.Context) {
	// Connection to the database
	db := c.MustGet("db").(*gorm.DB)
	// db := InitDb()
	// // Close connection database
	// defer db.Close()

	var rider models.Rider
	// SELECT * FROM riders
	db.Debug().Find(&rider)

	// Display JSON result
	c.JSON(200, rider)

	// curl -i http://localhost:8080/api/v1/users
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var rider models.Rider

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Debug().Set("gorm:auto_preload", true).Find(&rider, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, rider.Serialize())
}

// func read(c *gin.Context) {
// 	// Connection to the database
// 	// db := InitDb()
// 	// // Close connection database
// 	// defer db.Close()
// 	db := c.MustGet("db").(*gorm.DB)

// 	id := c.Params.ByName("id")
// 	//cid := c.Params.ByName("class_name_id")

// 	// cid := c.Params.ByName("card_id")
// 	var rider models.Rider
// 	db.debug().Set("gorm:auto_preload", true).Find(&rider, id)

// 	// Pre-Load Class Example - Preload works
// 	// Invoke-RestMethod -Uri "http://127.0.0.1:8080/api/v1/riders/3" -methode Get | ConvertFrom-Json
// 	// SELECT * FROM "riders"  WHERE ("riders"."id" = '3')
// 	// SELECT * FROM "classes"  WHERE ("id" IN ('2'))
// 	// db.Debug().Preload("Class").Find(&rider, id)

// 	// Test pre_load anabled for rider
// 	// db.Set("gorm:auto_preload", true).Find(&rider)
// 	db.Debug().Find(&rider, id)

// 	if rider.ID != 0 {
// 		// Display JSON result
// 		c.JSON(200, rider)
// 	} else {
// 		// Display JSON error
// 		c.JSON(404, gin.H{"error": "Rider not found"})
// 	}
// 	// Invoke-RestMethod -Uri "http://127.0.0.1:8080/api/v1/rider/1" -Method get | ConvertFrom-Json
// }

func update(c *gin.Context) {
	// // Connection to the database
	// db := InitDb()
	// // Close connection database
	// defer db.Close()
	db := c.MustGet("db").(*gorm.DB)

	// Get id user
	id := c.Params.ByName("id")
	var rider models.Rider
	// SELECT * FROM users WHERE id = 1;
	db.First(&rider, id)

	if rider.Firstname != "" && rider.Lastname != "" {

		if rider.ID != 0 {
			var newUser models.Rider
			c.Bind(&newUser)

			result := models.Rider{
				ID:        newUser.ID,
				Firstname: newUser.Firstname,
				Lastname:  newUser.Lastname,
				Birthday:  newUser.Birthday,
				Gender:    newUser.Gender,
				Mail:      newUser.Mail,
			}

			// UPDATE users SET firstname='newUser.Firstname', lastname='newUser.Lastname' WHERE id = user.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Rider not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

}

func remove(c *gin.Context) {
	// // Connection to the database
	// db := InitDb()
	// // Close connection database
	// defer db.Close()
	db := c.MustGet("db").(*gorm.DB)

	// Get id user
	id := c.Params.ByName("id")
	var rider models.Rider
	// SELECT * FROM users WHERE id = 1;
	db.First(&rider, id)

	if rider.ID != 0 {
		// DELETE FROM users WHERE id = user.Id
		db.Debug().Delete(&rider)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Rider #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Rider not found"})
	}
}
