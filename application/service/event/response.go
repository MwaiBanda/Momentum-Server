package event

type Response struct {
	Data []Event `json:"data" binding:"required"`
}

type Event struct {
	ID         string     `json:"id" binding:"required"`
	Attributes Attributes `json:"attributes" binding:"required"`
}

type Attributes struct {
	AllDayEvent           bool   `json:"all_day_event" binding:"required"`
	Description           string `json:"open_graph_description" binding:"required"`
	Name                  string `json:"name" binding:"required"`
	Cover                 string `json:"image_url" binding:"required"`
	StartTime             string `json:"starts_at" binding:"required"`
	EndsTime              string `json:"ends_at" binding:"required"`
	RecurrenceDescription string `json:"recurrence_description" binding:"required"`
	Recurring             bool   `json:"recurring" binding:"required"`
}
