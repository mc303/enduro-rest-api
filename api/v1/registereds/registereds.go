package registereds

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesClasses applies router to the gin Engine
func ApplyRoutesRegistereds(r *gin.RouterGroup) {
	registereds := r.Group("/registereds")
	{
		// classes.POST("/", create) //middlewares.Authorized,
		registereds.GET("/", list)
		registereds.GET("/:id", read)
		// classes.DELETE("/:id", remove) //middlewares.Authorized,
		// classes.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
