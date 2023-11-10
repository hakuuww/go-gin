package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hakuuww/go-gin/controllers"
	"log"
)

func main() {
	server := gin.Default()

	videocontroller := controllers.NewVideoController()

	videoGroup := server.Group("/videos")

	videoGroup.GET("/", videocontroller.GetAll)
	videoGroup.POST("/id", videocontroller.Update)
	videoGroup.PUT("/", videocontroller.Create)
	videoGroup.DELETE("/videos", videocontroller.Delete)

	log.Fatalln(server.Run("localhost:8080"))
}
