package event

type Response struct {
	Data []Event `json:"data"`
}

type Event struct {
	ID         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	AllDayEvent           bool   `json:"all_day_event"`
	Description           string `json:"open_graph_description"`
	Name                  string `json:"name"`
	Cover                 string `json:"image_url"`
	StartTime             string `json:"starts_at"`
	EndsTime              string `json:"ends_at"`
	RecurrenceDescription string `json:"recurrence_description"`
	Recurring             bool   `json:"recurring"`
}
