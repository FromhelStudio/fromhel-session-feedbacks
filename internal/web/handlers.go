package web

import (
	"context"
	"strconv"

	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/models"
	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateRatingHandler(c *gin.Context, uri string) {
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

	service, err := services.NewRatingService(uri, context.Background())

	if err != nil {
		panic(err)
	}

	if err := service.CreateRating(&rating); err != nil {
		handleError(c, 500, "Internal server error")
		return
	}

	c.JSON(201, gin.H{
		"message":    "Rating created",
		"statusCode": 201,
	})
}

func GetRatingsHandler(c *gin.Context, uri string) {
	service, err := services.NewRatingService(uri, context.Background())

	if err != nil {
		panic(err)
	}

	page := c.Param("page")
	if page == "" {
		page = "1"
	}

	// Convert page to int
	pageInt, err := strconv.ParseInt(page, 10, 64)

	if err != nil || pageInt < 1 {
		handleError(c, 400, "Invalid page")
		return
	}

	ratings, err := service.GetRatings(pageInt)

	if err != nil {
		handleError(c, 500, "Internal server error")
		return
	}

	c.JSON(200, gin.H{
		"ratings":    ratings,
		"statusCode": 200,
	})
}

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message, "statusCode": statusCode})
}
