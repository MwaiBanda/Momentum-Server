package model

type MessageResponse struct {
	Data []Message `json:"data"`
}

type Message struct {
	ID        string `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Series	string `json:"series"`
	Title     string `json:"title"`
	Preacher  string `json:"preacher"`	
	Date	  string `json:"date"`
	CreatedOn string `json:"createdOn"`
	Published bool   `json:"published"`
	HasOrder  bool   `json:"hasOrder"`
	Passages []PassageResponse `json:"passages"`
}

type PassageResponse struct {	
	ID        string `json:"id"`
	Header   string `json:"header"`
	Verse    string `json:"verse"`
	Message string `json:"message"`
	Type string `json:"type"`
	Order int `json:"order"`
	Notes []Note `json:"notes"`
}

type Note struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	UserID	string `json:"userId"`
}

type NoteRequest struct {
	PassageID        string `json:"passageId"`
	UserID        string `json:"userId"`
	Content   string `json:"content"`
}