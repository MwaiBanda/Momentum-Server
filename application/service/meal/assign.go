package meal

import (
	"context"
	"errors"

	domainmeal "Momentum/domain/meal"
)

func (service *Service) AssignVolunteer(ctx context.Context, input VolunteerAssignmentInput) (VolunteerAssignmentResponse, error) {
	if service == nil || service.repository == nil ||
		input.MealID == "" ||
		input.VolunteeredMeal.ID == "" ||
		input.VolunteeredMeal.User.ID == "" {
		return VolunteerAssignmentResponse{}, ErrInvalidInput
	}

	_, err := service.repository.AssignVolunteer(ctx, domainmeal.AssignVolunteerParams{
		MealID:            input.MealID,
		VolunteeredMealID: input.VolunteeredMeal.ID,
		UserID:            input.VolunteeredMeal.User.ID,
		Description:       input.VolunteeredMeal.Description,
		Notes:             input.VolunteeredMeal.Notes,
	})
	if err != nil {
		if errors.Is(err, domainmeal.ErrNotFound) {
			return VolunteerAssignmentResponse{}, ErrNotFound
		}
		if errors.Is(err, domainmeal.ErrVolunteeredMealNotFound) {
			return VolunteerAssignmentResponse{}, ErrVolunteeredMealNotFound
		}
		return VolunteerAssignmentResponse{}, err
	}

	return VolunteerAssignmentResponse{
		MealID: input.MealID,
		VolunteeredMeal: VolunteeredMealResponse{
			ID:          input.VolunteeredMeal.ID,
			CreatedOn:   input.VolunteeredMeal.CreatedOn,
			Description: input.VolunteeredMeal.Description,
			Notes:       input.VolunteeredMeal.Notes,
			Date:        input.VolunteeredMeal.Date,
			User:        UserResponse{ID: input.VolunteeredMeal.User.ID},
			Updates:     input.VolunteeredMeal.Updates,
		},
	}, service.invalidateCache(ctx)
}
