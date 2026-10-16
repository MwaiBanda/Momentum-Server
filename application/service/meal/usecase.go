package meal

import "context"

type UseCase interface {
	List(ctx context.Context) ([]Response, error)
	Create(ctx context.Context, input Input) (Input, error)
	Update(ctx context.Context, input Input) (Input, error)
	Delete(ctx context.Context, id string) (Response, error)
	AssignVolunteer(ctx context.Context, input VolunteerAssignmentInput) (VolunteerAssignmentInput, error)
	UpdateVolunteeredMeal(ctx context.Context, input UpdateVolunteeredMealInput) (UpdateVolunteeredMealInput, error)
}
