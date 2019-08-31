package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Heartbeat is simple keep-alive handler
func Heartbeat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
