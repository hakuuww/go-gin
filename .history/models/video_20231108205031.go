package models


type Video struct {
	Id int 
	Title string `json:"title"`
	Description string `json:description"`
	Author string `json:description"`
}