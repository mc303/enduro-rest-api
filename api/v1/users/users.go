package stages

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesUsers applies router to the gin Engine
func ApplyRoutesUsers(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		// stages.POST("/", create) //middlewares.Authorized,
		users.GET("/", list)
		users.GET("/:id", read)
		// stages.DELETE("/:id", remove) //middlewares.Authorized,
		// stages.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
