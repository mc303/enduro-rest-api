package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mc303/gin-rest-api-sample/api/v1/classes"
	"github.com/mc303/gin-rest-api-sample/api/v1/events"
	"github.com/mc303/gin-rest-api-sample/api/v1/riders"
	"github.com/mc303/gin-rest-api-sample/api/v1/runs"
	"github.com/mc303/gin-rest-api-sample/api/v1/stages"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		//auth.ApplyRoutes(v1)
		classes.ApplyRoutesClasses(v1)
		riders.ApplyRoutesRiders(v1)
		events.ApplyRoutesEvents(v1)
		runs.ApplyRoutesRuns(v1)
		stages.ApplyRoutesStages(v1)
	}
}
