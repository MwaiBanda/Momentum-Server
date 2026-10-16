package volunteer_service

import (
	"time"

	domainuser "Momentum/domain/user"
	domainservice "Momentum/domain/volunteer_service"
)

type ServiceResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Image string `json:"image"`
}

type ServiceDetailResponse struct {
	ID       string                     `json:"id"`
	Name     string                     `json:"name"`
	Type     string                     `json:"type"`
	Image    string                     `json:"image"`
	Services []VolunteerServiceResponse `json:"services"`
}

type VolunteerServiceResponse struct {
	ID          string                        `json:"id"`
	Organizer   string                        `json:"organizer"`
	Description string                        `json:"description"`
	Email       string                        `json:"email"`
	Phone       string                        `json:"phone"`
	Days        []VolunteerServiceDayResponse `json:"days"`
	UserID      string                        `json:"user_id"`
	ServiceID   string                        `json:"service_id"`
}

type VolunteerServiceDayResponse struct {
	ID    string       `json:"id"`
	Notes string       `json:"notes"`
	Date  string       `json:"date"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID        string        `json:"id"`
	Email     string        `json:"email"`
	Fullname  string        `json:"fullname"`
	Phone     string        `json:"phone"`
	CreatedOn string        `json:"createdOn"`
	Meals     []interface{} `json:"meals"`
}

type CreateInput struct {
	ID          string           `json:"id"`
	Organizer   string           `json:"organizer"`
	Description string           `json:"description"`
	Email       string           `json:"email"`
	Phone       string           `json:"phone"`
	Days        []CreateDayInput `json:"days"`
	UserID      string           `json:"user_id"`
	ServiceID   string           `json:"service_id"`
}

type CreateDayInput struct {
	Date string `json:"date"`
}

type CreateResponse = CreateInput

type UpdateDayInput struct {
	ID    string       `json:"id"`
	Notes string       `json:"notes"`
	Date  string       `json:"date"`
	User  UserResponse `json:"user"`
}

type UpdateDayResponse = UpdateDayInput

func toServiceResponse(service domainservice.Service) ServiceResponse {
	return ServiceResponse{
		ID:    service.ID,
		Name:  service.Name,
		Type:  service.Type,
		Image: service.Image,
	}
}

func toDetailResponse(detail domainservice.ServiceDetail) ServiceDetailResponse {
	response := ServiceDetailResponse{
		ID:       detail.Service.ID,
		Name:     detail.Service.Name,
		Type:     detail.Service.Type,
		Image:    detail.Service.Image,
		Services: make([]VolunteerServiceResponse, 0, len(detail.Services)),
	}

	for _, service := range detail.Services {
		days := make([]VolunteerServiceDayResponse, 0, len(service.Days))
		for _, day := range service.Days {
			days = append(days, VolunteerServiceDayResponse{
				ID:    day.ID,
				Notes: day.Notes,
				Date:  day.Date,
				User:  toUserResponse(day.User),
			})
		}

		response.Services = append(response.Services, VolunteerServiceResponse{
			ID:          service.ID,
			Organizer:   service.Organizer,
			Description: service.Description,
			Email:       service.Email,
			Phone:       service.Phone,
			Days:        days,
			UserID:      service.UserID,
			ServiceID:   service.ServiceID,
		})
	}

	return response
}

func toUserResponse(user *domainuser.User) UserResponse {
	if user == nil {
		return UserResponse{}
	}

	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Fullname:  user.Fullname,
		Phone:     user.Phone,
		CreatedOn: user.CreatedOn.Format(time.RFC3339Nano),
	}
}
