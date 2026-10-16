package meal

import (
	"context"
	"errors"
	"log"

	"Momentum/constants"
	domainmeal "Momentum/domain/meal"
)

func (service *Service) UpdateVolunteeredMeal(ctx context.Context, input UpdateVolunteeredMealInput) (UpdateVolunteeredMealInput, error) {
	if service == nil || service.repository == nil || input.ID == "" {
		return UpdateVolunteeredMealInput{}, ErrInvalidInput
	}

	_, err := service.repository.UpdateVolunteeredMeal(ctx, domainmeal.UpdateVolunteeredMealParams{
		ID:          input.ID,
		Description: input.Description,
		Notes:       input.Notes,
	})
	if err != nil {
		if errors.Is(err, domainmeal.ErrVolunteeredMealNotFound) {
			return UpdateVolunteeredMealInput{}, ErrVolunteeredMealNotFound
		}
		return UpdateVolunteeredMealInput{}, err
	}

	return input, service.invalidateCache(ctx)
}

func (service *Service) invalidateCache(ctx context.Context) error {
	if service.cacheStore == nil {
		return nil
	}
	if err := service.cacheStore.Delete(ctx, constants.MealsKey); err != nil {
		log.Printf("[meal-cache] delete %q: %v", constants.MealsKey, err)
	}
	return nil
}
