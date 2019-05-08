package races

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesRaces applies router to the gin Engine
func ApplyRoutesRaces(r *gin.RouterGroup) {
	races := r.Group("/races")
	{
		// events.POST("/", create) //middlewares.Authorized,
		races.GET("/", list)
		races.GET("/:id", read)
		// events.DELETE("/:id", remove) //middlewares.Authorized,
		// events.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
