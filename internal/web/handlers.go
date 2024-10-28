package web

import (
	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/models"
	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateRatingHandler(c *gin.Context, uri *string) {
	var ratingDTO models.RatingDTO
	if err := c.ShouldBindJSON(&ratingDTO); err != nil {
		handleError(c, 400, "Invalid request")
		return
	}

	if ratingDTO.Rating < 1 || ratingDTO.Rating > 5 {
		handleError(c, 400, "Rating must be between 1 and 5")
		return
	}

	rating := services.Rating{
		Rating:   ratingDTO.Rating,
		Feedback: ratingDTO.Feedback,
	}

	service := services.NewRatingService(uri)

	if err := service.CreateRating(rating); err != nil {
		handleError(c, 500, "Internal server error")
		return
	}

	c.JSON(201, gin.H{
		"message":    "Rating created",
		"statusCode": 201,
	})
}

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message, "statusCode": statusCode})
}
