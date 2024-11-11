package router

import (
	"log"
	"os"

	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Initialize() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}

	router := gin.Default()

	router.POST("/rating", func(c *gin.Context) {
		web.CreateRatingHandler(c, uri)
	})

	router.GET("/rating/:page", func(c *gin.Context) {
		web.GetRatingsHandler(c, uri)
	})

	router.POST("/session", func(ctx *gin.Context) {
		web.CreateSessionHandler(ctx, uri)
	})

	port, exists := os.LookupEnv("PORT")
	if !exists{
		port = "8080" 
	}
	router.Run(":" + port)

}
