package main 

import (
	"github.com/hakuuww/go-gin/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main(){
	server := gin.Default()

	videocontroller :=  controllers.NewVideoController()

	videoGroup:= server.Group("/videos")
	controllers.g 

	

	videoGroup.GET("/videos",videocontroller.GetAll)
	videoGroup.POST("/videos",videocontroller.Update)
	videoGroup.PUT("/videos",videocontroller.Create)
	videoGroup.DELETE("/videos",videocontroller.Delete)


	log.Fatalln(server.Run("localhost:8080"))
}