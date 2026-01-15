package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadinessHandler responds to readiness probe requests.
func ReadinessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
	})
}

// LivenessHandler responds to liveness probe requests.
func LivenessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
	})
}
