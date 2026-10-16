package message

import "time"

type Message struct {
	ID        string
	Thumbnail string
	Series    string
	Title     string
	Preacher  string
	Date      string
	CreatedOn time.Time
	Published bool
	HasOrder  bool
	Passages  []Passage
}

type Passage struct {
	ID        string
	Header    string
	Verse     string
	Message   string
	MessageID string
	CreatedOn time.Time
	Order     int
	Notes     []Note
}

type Note struct {
	ID        string
	Content   string
	UserID    string
	PassageID string
	CreatedOn time.Time
	UpdatedOn time.Time
}

type CreateParams struct {
	ID        string
	Thumbnail string
	Series    string
	Title     string
	Preacher  string
	Date      string
	Passages  []CreatePassageParams
}

type CreatePassageParams struct {
	ID      string
	Header  *string
	Verse   *string
	Message *string
	Order   int
}

type UpdateParams struct {
	ID        string
	Thumbnail string
	Series    string
	Title     string
	Preacher  string
	Date      string
	Published bool
	Passages  []UpdatePassageParams
}

type UpdatePassageParams struct {
	ID      string
	Header  string
	Verse   string
	Message string
}

type CreateNoteParams struct {
	ID        string
	Content   string
	UserID    string
	PassageID string
}

type UpdateNoteParams struct {
	ID      string
	Content string
}
