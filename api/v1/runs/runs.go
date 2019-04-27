package runs

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesRuns applies router to the gin Engine
func ApplyRoutesRuns(r *gin.RouterGroup) {
	runs := r.Group("/runs")
	{
		// runs.POST("/", create) //middlewares.Authorized,
		runs.GET("/", list)
		runs.GET("/:id", read)
		// runs.GET("/event/:id", readevent)
		// runs.GET("/stage/:id", readstage)
		// runs.GET("/rider/:id", readrider)
		// runs.DELETE("/:id", remove) //middlewares.Authorized,
		// runs.PATCH("/:id", update)  //middlewares.Authorized,

	}
}
