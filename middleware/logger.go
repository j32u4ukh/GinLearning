package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] | %d | %s %s\n",
			params.TimeStamp,
			params.Method,
			params.StatusCode,
			params.ClientIP,
			params.Path,
		)
	})
}
