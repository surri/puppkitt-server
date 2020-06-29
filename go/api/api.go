package api

import (
	"github.com/gin-gonic/gin"
	"puppkitt.com/v1/api/auth"
	"puppkitt.com/v1/api/posts"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/ping", ping)
		auth.ApplyRoutes(api)
		posts.ApplyRoutes(api)
	}
}
