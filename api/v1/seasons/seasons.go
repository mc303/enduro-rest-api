package seasons

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesClasses applies router to the gin Engine
func ApplyRoutesClasses(r *gin.RouterGroup) {
	seasons := r.Group("/seasons")
	{
		// classes.POST("/", create) //middlewares.Authorized,
		seasons.GET("/", list)
		seasons.GET("/:id", read)
		// classes.DELETE("/:id", remove) //middlewares.Authorized,
		// classes.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
