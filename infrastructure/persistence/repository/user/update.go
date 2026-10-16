package user

import (
	"context"
	"fmt"

	domainuser "Momentum/domain/user"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Update(ctx context.Context, params domainuser.UpdateParams) (domainuser.User, error) {
	var userEntity domainuser.User
	query := persistence.LoadFile("repository/user/sql/update.sql")
	if err := repository.db.GetContext(
		ctx,
		&userEntity,
		query,
		params.ID,
		params.Email,
		params.Fullname,
		params.Phone,
	); err != nil {
		return domainuser.User{}, fmt.Errorf("update user %q: %w", params.ID, err)
	}

	return userEntity, nil
}
