package models


type Video struct {
	Id int 
	Title string `json:"title"`
	Description string `json:descrip"`
	Author string
}