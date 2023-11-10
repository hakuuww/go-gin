package middlewares

import "github.com/gin-gonic/gin"

func MyLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter


}