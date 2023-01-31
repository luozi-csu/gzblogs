package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/utils/logx"
)

var (
	hostname, _ = os.Hostname()
)

func RequestLogger(c *gin.Context) {
	start := time.Now()

	defer func() {
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()

		msg := fmt.Sprintf("[%s %s] %d %v", c.Request.Method, c.Request.URL, statusCode, latency)
		requestInfo := fmt.Sprintf("msg=%s hostname=%s clientIP=%s userAgent=%s", msg, hostname, clientIP, clientUserAgent)

		if len(c.Errors) > 0 {
			logx.Errorf(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			if statusCode >= http.StatusInternalServerError {
				logx.Errorf(requestInfo)
			} else if statusCode >= http.StatusBadRequest {
				logx.Warnf(requestInfo)
			} else {
				logx.Infof(requestInfo)
			}
		}
	}()

	c.Next()
}
