package disciplines

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutesDisciplines applies router to the gin Engine
func ApplyRoutesDisciplines(r *gin.RouterGroup) {
	disciplines := r.Group("/disciplines")
	{
		// disciplines.POST("/", create) //middlewares.Authorized,
		disciplines.GET("/", list)
		disciplines.GET("/:id", read)
		// disciplines.DELETE("/:id", remove) //middlewares.Authorized,
		// disciplines.PATCH("/:id", update)  //middlewares.Authorized,
	}
}
