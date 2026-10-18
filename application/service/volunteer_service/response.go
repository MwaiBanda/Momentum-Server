package volunteer_service

import (
	"time"

	domainuser "Momentum/domain/user"
	domainservice "Momentum/domain/volunteer_service"
)

type ServiceResponse struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Type  string `json:"type" binding:"required"`
	Image string `json:"image" binding:"required"`
}

type ServiceDetailResponse struct {
	ID       string                     `json:"id" binding:"required"`
	Name     string                     `json:"name" binding:"required"`
	Type     string                     `json:"type" binding:"required"`
	Image    string                     `json:"image" binding:"required"`
	Services []VolunteerServiceResponse `json:"services" binding:"required"`
}

type VolunteerServiceResponse struct {
	ID          string                        `json:"id" binding:"required"`
	Organizer   string                        `json:"organizer" binding:"required"`
	Description string                        `json:"description" binding:"required"`
	Email       string                        `json:"email" binding:"required"`
	Phone       string                        `json:"phone" binding:"required"`
	Days        []VolunteerServiceDayResponse `json:"days" binding:"required"`
	UserID      string                        `json:"user_id" binding:"required"`
	ServiceID   string                        `json:"service_id" binding:"required"`
}

type VolunteerServiceDayResponse struct {
	ID    string       `json:"id" binding:"required"`
	Notes string       `json:"notes" binding:"required"`
	Date  string       `json:"date" binding:"required"`
	User  UserResponse `json:"user" binding:"required"`
}

type UserResponse struct {
	ID        string        `json:"id" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Fullname  string        `json:"fullname" binding:"required"`
	Phone     string        `json:"phone" binding:"required"`
	CreatedOn string        `json:"createdOn" binding:"required"`
	Meals     []interface{} `json:"meals" binding:"required"`
}

type UserInput struct {
	ID string `json:"id" binding:"required"`
}

type CreateInput struct {
	ID          string           `json:"id"`
	Organizer   string           `json:"organizer" binding:"required"`
	Description string           `json:"description" binding:"required"`
	Email       string           `json:"email" binding:"required"`
	Phone       string           `json:"phone" binding:"required"`
	Days        []CreateDayInput `json:"days" binding:"required"`
	UserID      string           `json:"user_id" binding:"required"`
	ServiceID   string           `json:"service_id" binding:"required"`
}

type CreateDayInput struct {
	Date string `json:"date" binding:"required"`
}

type CreateResponse struct {
	ID          string           `json:"id" binding:"required"`
	Organizer   string           `json:"organizer" binding:"required"`
	Description string           `json:"description" binding:"required"`
	Email       string           `json:"email" binding:"required"`
	Phone       string           `json:"phone" binding:"required"`
	Days        []CreateDayInput `json:"days" binding:"required"`
	UserID      string           `json:"user_id" binding:"required"`
	ServiceID   string           `json:"service_id" binding:"required"`
}

type UpdateDayInput struct {
	ID    string    `json:"id" binding:"required"`
	Notes string    `json:"notes" binding:"required"`
	Date  string    `json:"date" binding:"required"`
	User  UserInput `json:"user" binding:"required"`
}

type UpdateDayResponse struct {
	ID    string       `json:"id" binding:"required"`
	Notes string       `json:"notes" binding:"required"`
	Date  string       `json:"date" binding:"required"`
	User  UserResponse `json:"user" binding:"required"`
}

func toCreateResponse(input CreateInput) CreateResponse {
	return CreateResponse{
		ID:          input.ID,
		Organizer:   input.Organizer,
		Description: input.Description,
		Email:       input.Email,
		Phone:       input.Phone,
		Days:        input.Days,
		UserID:      input.UserID,
		ServiceID:   input.ServiceID,
	}
}

func toUpdateDayResponse(input UpdateDayInput) UpdateDayResponse {
	return UpdateDayResponse{
		ID:    input.ID,
		Notes: input.Notes,
		Date:  input.Date,
		User:  UserResponse{ID: input.User.ID},
	}
}

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
