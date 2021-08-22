package middleware

import (
	"github.com/LearnGin/middleware/debug"
	"github.com/gin-gonic/gin"
)

func RegisterMiddleware(r *gin.Engine) {
	r.Use(debug.DebugMiddleWare())
}
