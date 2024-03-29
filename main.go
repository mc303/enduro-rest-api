package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mc303/gin-rest-api-sample/api"
	"github.com/mc303/gin-rest-api-sample/database"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	//app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + port)  // listen to given port
}
