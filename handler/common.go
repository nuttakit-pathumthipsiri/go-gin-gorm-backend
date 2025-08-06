package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleErrorResponse is a helper function to handle error responses consistently
func handleErrorResponse(c *gin.Context, err error) {
	switch err.Error() {
	case "topic not found", "topic detail not found":
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case "topic name already exists", "topic detail name already exists":
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case "order number already exists":
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order number already exists"})
	case "invalid topic ID format":
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Topic ID format"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
