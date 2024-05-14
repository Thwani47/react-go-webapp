package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func writeInternalServerError(c *gin.Context, message string, err error) {
	if message != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func writeBadRequestError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
