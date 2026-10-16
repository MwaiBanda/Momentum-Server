package volunteer_service

import "context"

type Repository interface {
	ListServices(ctx context.Context) ([]Service, error)
	GetByType(ctx context.Context, serviceType string) (ServiceDetail, error)
	Create(ctx context.Context, params CreateParams) error
	UpdateDay(ctx context.Context, params UpdateDayParams) (VolunteerServiceDay, error)
}
