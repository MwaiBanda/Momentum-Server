package model

type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type		string `json:"type"`
	Image       string `json:"image"`
}	

type ServiceResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type		string `json:"type"`
	Image       string `json:"image"`
	Services   []VolunteerService `json:"services"`
}	

type VolunteerService struct{
	Id			  string  `json:"id"`
	Organizer     string  `json:"organizer"`
	Description   string  `json:"description"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	Days          []VolunteerServiceDay `json:"days"`
	UserId        string  `json:"user_id"`
	ServiceId     string  `json:"service_id"`
}
	
type VolunteerServiceDay struct {
	Id		  string `json:"id"`
	Notes     string `json:"notes"`
	Date       string `json:"date"`
}
  
type VolunteerServiceRequest struct {
	Id			  	string                   `json:"id"`
	Organizer       string                   `json:"organizer"`
	Description          string                   `json:"description"`
	Email           string                   `json:"email"`
	Phone           string                   `json:"phone"`
	Days           []VolunteerServiceDayRequest `json:"days"`
	UserId          string                   `json:"user_id"`
	ServiceId       string                   `json:"service_id"`
}

type VolunteerServiceDayRequest struct {
	Date        string `json:"date"`
}
