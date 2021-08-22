package debug

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func DebugMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		fmt.Printf("[Debug Middleware] - Request: %v\n", req)
		c.Next()
		fmt.Printf("[Debug Middleware] - Response: %v\n", c.Writer.Status())
	}
}
