package models


//If the frontend uses form to submitt 
type Video struct {
	Id int 
	Title string `json:"title"`
	Description string `json:"description"`
	Author string `json:"author"`
}



