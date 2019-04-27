package stages

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesStages applies router to the gin Engine
func ApplyRoutesStages(r *gin.RouterGroup) {
	stages := r.Group("/stages")
	{
		// stages.POST("/", create) //middlewares.Authorized,
		stages.GET("/", list)
		stages.GET("/:id", read)
		// stages.DELETE("/:id", remove) //middlewares.Authorized,
		// stages.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
