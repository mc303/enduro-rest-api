package events

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesEvents applies router to the gin Engine
func ApplyRoutesEvents(r *gin.RouterGroup) {
	events := r.Group("/events")
	{
		// events.POST("/", create) //middlewares.Authorized,
		events.GET("/", list)
		events.GET("/:id", read)
		// events.DELETE("/:id", remove) //middlewares.Authorized,
		// events.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
