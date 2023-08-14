package model

type Participant struct {
	Id   string       `json:"id"`
	User UserResponse `json:"user"`
}

type ParticipantRequest struct {
	UserId string `json:"user_id"`
	MealId string `json:"meal_id"`
}
