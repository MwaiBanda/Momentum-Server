package model

import "time"

type SermonResponse struct {
	CurrentPage  int      `json:"current_page"`
	Data         []Sermon `json:"data"`
	FirstPageURL string   `json:"first_page_url"`
	From         int      `json:"from"`
	LastPage     int      `json:"last_page"`
	LastPageURL  string   `json:"last_page_url"`
	Links        []struct {
		URL    any    `json:"url"`
		Label  string `json:"label"`
		Active bool   `json:"active"`
	} `json:"links"`
	NextPageURL string `json:"next_page_url"`
	Path        string `json:"path"`
	PerPage     int    `json:"per_page"`
	PrevPageURL any    `json:"prev_page_url"`
	To          int    `json:"to"`
	Total       string `json:"total"`
}

type Sermon struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	DetailURL string `json:"detail_url"`
	Date      struct {
		Carbon time.Time `json:"carbon"`
		Date   string    `json:"date"`
	} `json:"date"`
	Preacher        string `json:"preacher"`
	Summary         any    `json:"summary"`
	Text            any    `json:"text"`
	InteractiveNote any    `json:"interactive_note"`
	Passages        []any  `json:"passages"`
	Series          struct {
		Title       string `json:"title"`
		Slug        string `json:"slug"`
		Description any    `json:"description"`
		Image       any    `json:"image"`
	} `json:"series"`
	Categories any `json:"categories"`
	Keywords   any `json:"keywords"`
	Media      struct {
		Image          any    `json:"image"`
		Audio          any    `json:"audio"`
		Video          string `json:"video"`
		VideoThumbnail string `json:"video_thumbnail"`
		Embed          any    `json:"embed"`
		Notes          any    `json:"notes"`
	} `json:"media"`
	Interactions struct {
		Shares int `json:"shares"`
		Likes  int `json:"likes"`
	} `json:"interactions"`
}
