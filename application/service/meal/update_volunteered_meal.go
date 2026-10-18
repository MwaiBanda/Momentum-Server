package meal

import (
	"context"
	"errors"
	"log"

	"Momentum/application/telemetry"
	domainmeal "Momentum/domain/meal"
)

func (service *Service) UpdateVolunteeredMeal(ctx context.Context, input UpdateVolunteeredMealInput) (UpdateVolunteeredMealResponse, error) {
	ctx, span := telemetry.Start(ctx, "application.meal.update_volunteered_meal")
	defer span.End()

	if service == nil || service.repository == nil || input.ID == "" {
		return UpdateVolunteeredMealResponse{}, ErrInvalidInput
	}

	_, err := service.repository.UpdateVolunteeredMeal(ctx, domainmeal.UpdateVolunteeredMealParams{
		ID:          input.ID,
		Description: input.Description,
		Notes:       input.Notes,
	})
	if err != nil {
		if errors.Is(err, domainmeal.ErrVolunteeredMealNotFound) {
			return UpdateVolunteeredMealResponse{}, ErrVolunteeredMealNotFound
		}
		return UpdateVolunteeredMealResponse{}, err
	}

	return UpdateVolunteeredMealResponse{
		ID:          input.ID,
		Description: input.Description,
		Notes:       input.Notes,
	}, service.invalidateCache(ctx)
}

func (service *Service) invalidateCache(ctx context.Context) error {
	if service.cacheStore == nil {
		return nil
	}
	if err := service.cacheStore.Delete(ctx, mealsCacheKey); err != nil {
		log.Printf("[meal-cache] delete %q: %v", mealsCacheKey, err)
	}
	return nil
}
