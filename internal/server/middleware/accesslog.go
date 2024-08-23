package middleware

import (
	"time"

	"github.com/dsc-sgu/shawty/internal/log"

	"github.com/gin-gonic/gin"
)

// Middleware to log every processed request.
func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		logger := log.S.With(
			"path", c.Request.URL.Path,
		)
		logger.Info("Got HTTP request")

		c.Next()

		status := c.Writer.Status()
		logger.Infow(
			"Processed HTTP request",
			"status_code", status,
			"time_elapsed", time.Since(startTime),
		)
	}
}
