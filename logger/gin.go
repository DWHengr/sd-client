package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// LoggerFunc gin logger
func LoggerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		// Stop timer
		Logger.Infow("[GIN]",
			zap.Float64("latency", time.Now().Sub(start).Seconds()),
			zap.String("clientIP", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.Int("statusCode", c.Writer.Status()),
			zap.String("errorMessage", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Int("bodySize", c.Writer.Size()),
			zap.String("path", path),
		)
	}
}
