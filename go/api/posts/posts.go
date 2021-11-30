package posts

import (
	"github.com/gin-gonic/gin"
	"hotpler.com/v1/lib/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	{
		posts.POST("/", middlewares.Authorized, create)
		posts.GET("/", list)
		posts.GET("/:id", read)
		posts.DELETE("/:id", middlewares.Authorized, remove)
		posts.PATCH("/:id", middlewares.Authorized, update)
	}
}
