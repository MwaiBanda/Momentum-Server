package meal

import "context"

type Repository interface {
	List(ctx context.Context) ([]Meal, error)
	Create(ctx context.Context, params CreateParams) (Meal, error)
	Update(ctx context.Context, params UpdateParams) (Meal, error)
	Delete(ctx context.Context, id string) (Meal, error)
	AssignVolunteer(ctx context.Context, params AssignVolunteerParams) (Meal, error)
	UpdateVolunteeredMeal(ctx context.Context, params UpdateVolunteeredMealParams) (VolunteeredMeal, error)
}
