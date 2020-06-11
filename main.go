package main

import(
	"github.com/gin-gonic/gin"
	// "os"
	// "fmt"
	// "log"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/someGet", someMethod)
	r.POST("/somePost", someMethod)
	r.PUT("/somePut", someMethod)
	r.DELETE("/someDelete", someMethod)
	r.PATCH("/somePatch", someMethod)
	r.HEAD("/someHead", someMethod)
	r.OPTIONS("/someOptins", someMethod)

	return r
}

func someMethod(c *gin.Context) {
	httpMethod := c.Request.Method
	c.JSON(200, gin.H{"status": "good", "sending": httpMethod})
}

func main() {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode()

	r := setupRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}