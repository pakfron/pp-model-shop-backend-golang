package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PathNotFound() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "path not found"})
	}

}
