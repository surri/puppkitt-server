package api

import (
	"github.com/gin-gonic/gin"
	"hotpler.com/v1/api/auth"
	"hotpler.com/v1/api/posts"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", ping)
		auth.ApplyRoutes(api)
		posts.ApplyRoutes(api)
	}
}
