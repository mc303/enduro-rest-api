package classes

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesClasses applies router to the gin Engine
func ApplyRoutesClasses(r *gin.RouterGroup) {
	classes := r.Group("/classes")
	{
		// classes.POST("/", create) //middlewares.Authorized,
		classes.GET("/", list)
		classes.GET("/:id", read)
		// classes.DELETE("/:id", remove) //middlewares.Authorized,
		// classes.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
