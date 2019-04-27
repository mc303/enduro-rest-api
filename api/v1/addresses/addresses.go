package addresses

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesAdresses applies router to the gin Engine
func ApplyRoutesAddresses(r *gin.RouterGroup) {
	addresses := r.Group("/addresses")
	{
		// events.POST("/", create) //middlewares.Authorized,
		addresses.GET("/", list)
		addresses.GET("/:id", read)
		// events.DELETE("/:id", remove) //middlewares.Authorized,
		// events.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
