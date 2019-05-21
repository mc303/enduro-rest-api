package results

import (
	"github.com/gin-gonic/gin"
	"github.com/mc303/gin-rest-api-sample/lib/middlewares"
)

// ApplyRoutesResults applies router to the gin Engine
func ApplyRoutesResults(r *gin.RouterGroup) {
	results := r.Group("/results")
	{
		// classes.POST("/", create) //middlewares.Authorized,
		results.GET("/", middlewares.Authorized, list)
		results.GET("/:id", middlewares.Authorized, read)
		// classes.DELETE("/:id", remove) //middlewares.Authorized,
		// classes.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
