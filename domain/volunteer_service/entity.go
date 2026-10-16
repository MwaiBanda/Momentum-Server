package volunteer_service

import (
	"time"

	domainuser "Momentum/domain/user"
)

type Service struct {
	ID        string
	Name      string
	Type      string
	Image     string
	CreatedOn time.Time
}

type VolunteerService struct {
	ID          string
	Description string
	Email       string
	Phone       string
	Organizer   string
	ServiceID   string
	UserID      string
	CreatedOn   time.Time
	Days        []VolunteerServiceDay
}

type VolunteerServiceDay struct {
	ID        string
	Notes     string
	Date      string
	UserID    string
	CreatedOn time.Time
	User      *domainuser.User
}

type ServiceDetail struct {
	Service  Service
	Services []VolunteerService
}

type CreateParams struct {
	ID          string
	Description string
	Email       string
	Phone       string
	Organizer   string
	ServiceID   string
	UserID      string
	Days        []CreateDayParams
}

type CreateDayParams struct {
	ID   string
	Date string
}

type UpdateDayParams struct {
	ID     string
	Notes  string
	UserID string
}
