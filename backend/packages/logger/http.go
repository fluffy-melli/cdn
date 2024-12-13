package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		ip := c.ClientIP()
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		fmt.Println("-  " + Fg_BrightBlue + ShortenString(ip, 20) + Fg_BrightGreen + ShortenString(method, 6) + ShortenString(fmt.Sprintf("%d", status), 6) + Fg_BrightYellow + ShortenString(path, 23) + "  " + Fg_BrightCyan + TimeToLargestUnit(duration.Nanoseconds()) + Reset)
	}
}
