package meal

import (
	"context"
	"errors"

	domainmeal "Momentum/domain/meal"
)

func (service *Service) AssignVolunteer(ctx context.Context, input VolunteerAssignmentInput) (VolunteerAssignmentInput, error) {
	if service == nil || service.repository == nil ||
		input.MealID == "" ||
		input.VolunteeredMeal.ID == "" ||
		input.VolunteeredMeal.User.ID == "" {
		return VolunteerAssignmentInput{}, ErrInvalidInput
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
			return VolunteerAssignmentInput{}, ErrNotFound
		}
		if errors.Is(err, domainmeal.ErrVolunteeredMealNotFound) {
			return VolunteerAssignmentInput{}, ErrVolunteeredMealNotFound
		}
		return VolunteerAssignmentInput{}, err
	}

	return input, service.invalidateCache(ctx)
}
