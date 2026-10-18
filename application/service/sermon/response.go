package sermon

import "time"

type Response struct {
	CurrentPage  int      `json:"current_page" binding:"required"`
	Data         []Sermon `json:"data" binding:"required"`
	FirstPageURL string   `json:"first_page_url" binding:"required"`
	From         int      `json:"from" binding:"required"`
	LastPage     int      `json:"last_page" binding:"required"`
	LastPageURL  string   `json:"last_page_url" binding:"required"`
	Links        []struct {
		URL    any    `json:"url" binding:"required"`
		Label  string `json:"label" binding:"required"`
		Active bool   `json:"active" binding:"required"`
	} `json:"links" binding:"required"`
	NextPageURL string `json:"next_page_url" binding:"required"`
	Path        string `json:"path" binding:"required"`
	PerPage     int    `json:"per_page" binding:"required"`
	PrevPageURL any    `json:"prev_page_url" binding:"required"`
	To          int    `json:"to" binding:"required"`
	Total       string `json:"total" binding:"required"`
}

type Sermon struct {
	ID        string `json:"id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Slug      string `json:"slug" binding:"required"`
	DetailURL string `json:"detail_url" binding:"required"`
	Date      struct {
		Carbon time.Time `json:"carbon" binding:"required"`
		Date   string    `json:"date" binding:"required"`
	} `json:"date" binding:"required"`
	Preacher        string `json:"preacher" binding:"required"`
	Summary         any    `json:"summary" binding:"required"`
	Text            any    `json:"text" binding:"required"`
	InteractiveNote any    `json:"interactive_note" binding:"required"`
	Passages        []any  `json:"passages" binding:"required"`
	Series          struct {
		Title       string `json:"title" binding:"required"`
		Slug        string `json:"slug" binding:"required"`
		Description any    `json:"description" binding:"required"`
		Image       any    `json:"image" binding:"required"`
	} `json:"series" binding:"required"`
	Categories any `json:"categories" binding:"required"`
	Keywords   any `json:"keywords" binding:"required"`
	Media      struct {
		Image          any    `json:"image" binding:"required"`
		Audio          any    `json:"audio" binding:"required"`
		Video          string `json:"video" binding:"required"`
		VideoThumbnail string `json:"video_thumbnail" binding:"required"`
		Embed          any    `json:"embed" binding:"required"`
		Notes          any    `json:"notes" binding:"required"`
	} `json:"media" binding:"required"`
	Interactions struct {
		Shares int `json:"shares" binding:"required"`
		Likes  int `json:"likes" binding:"required"`
	} `json:"interactions" binding:"required"`
}
