package middlewares

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func MyLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string{
		return fmt.Print("%s,%s,%s", params.ClientIP,params.Method)

	})

}