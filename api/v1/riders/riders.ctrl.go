package riders

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mc303/enduro-rest-api/database/models"
	"github.com/mc303/enduro-rest-api/lib/common"
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
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var riders []models.Rider

	// db.Preload("Stages").Preload("TypeOfRaces").Find(&event)
	// db.Model(&events).Related(&card)
	if cursor == "" {
		// db.Model(&rider).Related(&run).Related(&result).Related(&address).Related(&registered)
		// db.Preload("Runs").Preload("Results").Preload("Addresses").Preload("Registereds")
		if err := db.Find(&riders).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		// db.Preload("Stages").Preload("TypeOfRaces").Where(condition, cursor)
		if err := db.Where(condition, cursor).Find(&riders).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(riders)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = riders[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var rider models.Rider
	// var run models.Run
	// var result models.Result
	// var address models.Address
	// var registered models.Registered

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	// db.Set("gorm:auto_preload", true)

	if err := db.Find(&rider, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, rider.Serialize())
}

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
		db.Delete(&rider)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Rider #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Rider not found"})
	}
}
