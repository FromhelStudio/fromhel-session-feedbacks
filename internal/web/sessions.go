package web

import (
	"context"

	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/models"
	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateSessionHandler(c *gin.Context, uri string) {
	var session models.SessionsDTO
	if err := c.ShouldBindJSON(&session); err != nil {
		handleError(c, 400, "Invalid request")
		return
	}

	isNotNil := allFieldsNonNil(session)

	if !isNotNil {
		handleError(c, 400, "Some fields are null")
		return
	}

	service, err := services.NewSessionService(uri, context.Background())

	if err != nil {
		handleError(c, 500, "Internal server error")
		return
	}

	err = service.CreateSession(session)

	if err != nil {
		handleError(c, 500, "Not saved")
		return
	}

	c.JSON(201, gin.H{
		"message":    "Created.",
		"statusCode": 201,
	})
}
