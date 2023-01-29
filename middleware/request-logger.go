package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/utils/logx"
)

func RequestLogger(logger *logx.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		defer func() {
			latency := time.Since(start)
			statusCode := c.Writer.Status()

			if len(c.Errors) > 0 {
				logger.Errorf(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				msg := fmt.Sprintf("[%s %s] %d %v", c.Request.Method, c.Request.URL, statusCode, latency)
				if statusCode >= http.StatusInternalServerError {
					logger.Errorf(msg)
				} else if statusCode >= http.StatusBadRequest {
					logger.Warnf(msg)
				} else {
					logger.Infof(msg)
				}
			}
		}()

		c.Next()
	}
}
