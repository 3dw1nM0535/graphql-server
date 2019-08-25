package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is smple keep-alive ping handler
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
