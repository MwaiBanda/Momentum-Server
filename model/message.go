package model

type MessageResponse struct {
	ID        string `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Series	string `json:"series"`
	Title     string `json:"title"`
	Preacher  string `json:"preacher"`	
	Date	  string `json:"date"`
	CreatedOn string `json:"createdOn"`
	Passages []PassageResponse `json:"passages"`
}

type PassageResponse struct {	
	ID        string `json:"id"`
	Header   string `json:"header"`
	Verse    string `json:"verse"`
	Message string `json:"message"`
}