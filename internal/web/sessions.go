package web

import (
	"context"
	"strconv"
	"strings"

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

func GetSessionsHandler(c *gin.Context, uri string) {
	service, err := services.NewSessionService(uri, context.Background())

	gameName := c.Query("game")

	if gameName == "" {
		handleError(c, 400, "Game is required")
		return
	}

	gameName = strings.Trim(strings.ToLower(gameName), " ")

	if gameName != "bulletspeel" && gameName != "cordel" {
		handleError(c, 400, "Invalid game")
		return
	}

	if err != nil {
		panic(err)
	}

	page := c.Param("page")
	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)

	if err != nil || pageInt < 1 {
		handleError(c, 400, "Invalid page")
		return
	}

	sessions, err := service.GetSession(gameName, pageInt)

	if err != nil {
		handleError(c, 500, "Internal server error")
		return
	}

	c.JSON(200, gin.H{
		"sessions":   sessions,
		"statusCode": 200,
	})
}
