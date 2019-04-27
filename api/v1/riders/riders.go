package riders

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesRiders applies router to the gin Engine
func ApplyRoutesRiders(r *gin.RouterGroup) {
	riders := r.Group("/riders")
	{
		riders.POST("/", create) //middlewares.Authorized,
		riders.GET("/", list)
		riders.GET("/:id", read)
		riders.DELETE("/:id", remove) //middlewares.Authorized,
		riders.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
