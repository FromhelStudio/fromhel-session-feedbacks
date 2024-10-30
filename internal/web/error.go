package web

import "github.com/gin-gonic/gin"

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message, "statusCode": statusCode})
}
