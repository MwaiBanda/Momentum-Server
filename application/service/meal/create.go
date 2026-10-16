package meal

import (
	"context"

	domainmeal "Momentum/domain/meal"

	"github.com/lucsky/cuid"
)

func (service *Service) Create(ctx context.Context, input Input) (Input, error) {
	if service == nil || service.repository == nil {
		return Input{}, ErrInvalidInput
	}

	_, err := service.repository.Create(ctx, toCreateParams(input))
	if err != nil {
		return Input{}, err
	}

	return input, service.invalidateCache(ctx)
}

func toCreateParams(input Input) domainmeal.CreateParams {
	volunteeredMeals := make([]domainmeal.CreateVolunteeredMealParams, len(input.Meals))
	for index, inputMeal := range input.Meals {
		volunteeredMeals[index] = domainmeal.CreateVolunteeredMealParams{
			ID:   cuid.New(),
			Date: inputMeal.Date,
		}
	}

	return domainmeal.CreateParams{
		ID:              cuid.New(),
		Recipient:       input.Recipient,
		Reason:          input.Reason,
		Email:           input.Email,
		Phone:           input.Phone,
		City:            input.City,
		NumOfAdults:     input.NumOfAdults,
		NumberOfKids:    input.NumberOfKids,
		Street:          input.Street,
		PreferredTime:   input.PreferredTime,
		Favourites:      input.Favourites,
		LeastFavourites: input.LeastFavourites,
		Allergies:       input.Allergies,
		Instructions:    input.Instructions,
		UserID:          input.UserID,
		Meals:           volunteeredMeals,
	}
}
