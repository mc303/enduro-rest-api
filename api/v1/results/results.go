package results

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesClasses applies router to the gin Engine
func ApplyRoutesResults(r *gin.RouterGroup) {
	results := r.Group("/results")
	{
		// classes.POST("/", create) //middlewares.Authorized,
		results.GET("/", list)
		results.GET("/:id", read)
		// classes.DELETE("/:id", remove) //middlewares.Authorized,
		// classes.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
