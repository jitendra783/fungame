package api

import (
	"fungame/pkg/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Route(obj service.ServiceLayer, logger *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(customLogger(logger))
	router.GET("/health", obj.Status)
	return router
}
func customLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		if c.FullPath() != "/health" {
			latency := time.Since(start).Milliseconds()
			// userID := c.GetString(config.USERID)
			// uID := c.GetString(config.REQUESTID)
			logger.Info(path,
				// zap.String("requestID", uID),
				// zap.String("userId", userID),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Int64("latency", latency),
			)
		}
	}

}
