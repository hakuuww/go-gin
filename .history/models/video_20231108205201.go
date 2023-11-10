package models


//If the frontend uses form to submitt requests, we can add form tags inside the struct
type Video struct {
	Id int 
	Title string `json:"title"`
	Description string `json:"description"`
	Author string `json:"author"`
}



