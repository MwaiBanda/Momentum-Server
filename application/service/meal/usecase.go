package meal

import "context"

type UseCase interface {
	List(ctx context.Context) ([]Response, error)
	Create(ctx context.Context, input Input) (CreateResponse, error)
	Update(ctx context.Context, input Input) (UpdateResponse, error)
	Delete(ctx context.Context, id string) (Response, error)
	AssignVolunteer(ctx context.Context, input VolunteerAssignmentInput) (VolunteerAssignmentResponse, error)
	UpdateVolunteeredMeal(ctx context.Context, input UpdateVolunteeredMealInput) (UpdateVolunteeredMealResponse, error)
}
