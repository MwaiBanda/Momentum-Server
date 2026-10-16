package user

import (
	"context"
	"fmt"

	domainuser "Momentum/domain/user"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Create(ctx context.Context, params domainuser.CreateParams) (domainuser.User, error) {
	var userEntity domainuser.User
	query := persistence.LoadFile("repository/user/sql/create.sql")
	if err := repository.db.GetContext(
		ctx,
		&userEntity,
		query,
		params.ID,
		params.Email,
		params.Fullname,
		params.Phone,
	); err != nil {
		return domainuser.User{}, fmt.Errorf("create user %q: %w", params.ID, err)
	}

	return userEntity, nil
}
