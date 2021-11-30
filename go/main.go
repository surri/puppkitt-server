package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hotpler.com/v1/api"
	"hotpler.com/v1/database"
	"hotpler.com/v1/lib/middlewares"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic("panic main")
		panic(err)
	}

	// initializes database
	db, _ := database.Initialize()
	gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + port)  // listen to given port
}
