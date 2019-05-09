package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mc303/enduro-rest-api/api/v1"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1.ApplyRoutes(api)
	}
}
