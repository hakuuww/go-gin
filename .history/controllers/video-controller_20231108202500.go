package controllers

import (
	"github.com/hakuuww/go-gin/models"
	"github.com/gin-gonic/gin"
)


type VideoController interface {
	GetAll (context *gin.Context)
	Update (context *gin.Context)
	Create (context *gin.Context)
	Delete (context *gin.Context)
}

type controller struct {
	videos []models.Video

}

func NewVideoController() VideoController {
	return &controller{}
}

func (c controller) GetAll (context *gin.Context){
	context.JSON(200, c.videos )
}

func (c *controller) Update (context *gin.Context){
	var newVideo models.Video

	//ShouldBind checks the Method and Content-Type to select a binding engine automatically, Depending on the "Content-Type" header different bindings are used, for example:
	// "application/json" --> JSON binding
	// "application/xml"  --> XML binding
	// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input. 
	// It decodes the json payload into the struct specified as a pointer. 
	// Like c.Bind() but this method does not set the response status code to 400 or abort if input is not valid.
	if err:= context.ShouldBind(&newVideo); err!= nil {
		context.String(400, "bad request, format does not match")
	}


	//In a Go for range loop, the values returned in each iteration are passed by value, not by pointer. 
	//This means that the video variable in your example is a copy of the value in the c.videos slice for each iteration.
	//Therefore in order to modify a specific value in the slice, we need to use the index to access it 
	for idx, video := range c.videos {
		if newVideo.Id == video.Id{
			c.videos[idx] = newVideo
			context.String(200, "video with id %d has been updated to the db",newVideo,)
		}  
	}



	//c.videos = append(c.videos,newVideo)
}

func (c controller) Create (context *gin.Context){

}

func (c controller) Delete (context *gin.Context){

}

