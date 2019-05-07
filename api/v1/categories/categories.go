package categories

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutescategories applies router to the gin Engine
func ApplyRoutesCategories(r *gin.RouterGroup) {
	categories := r.Group("/categories")
	{
		// categories.POST("/", create) //middlewares.Authorized,
		categories.GET("/", list)
		categories.GET("/:id", read)
		// categories.DELETE("/:id", remove) //middlewares.Authorized,
		// categories.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
