package meal

import (
	"context"

	domainmeal "Momentum/domain/meal"
)

func (service *Service) Update(ctx context.Context, input Input) (UpdateResponse, error) {
	if service == nil || service.repository == nil || input.ID == "" {
		return UpdateResponse{}, ErrInvalidInput
	}

	_, err := service.repository.Update(ctx, domainmeal.UpdateParams{
		ID:              input.ID,
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
	})
	if err != nil {
		return UpdateResponse{}, err
	}

	return UpdateResponse{
		ID:              input.ID,
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
		Meals:           input.Meals,
		UserID:          input.UserID,
	}, service.invalidateCache(ctx)
}
