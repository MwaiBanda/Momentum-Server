package volunteer_service

import "context"

type UseCase interface {
	ListServices(ctx context.Context) ([]ServiceResponse, error)
	GetByType(ctx context.Context, serviceType string) (ServiceDetailResponse, error)
	Create(ctx context.Context, input CreateInput) (CreateResponse, error)
	UpdateDay(ctx context.Context, input UpdateDayInput) (UpdateDayResponse, error)
}
