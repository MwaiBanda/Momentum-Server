package user

import (
	"context"

	domainuser "Momentum/domain/user"
	"Momentum/infrastructure/persistence"
	"fmt"
)

func (repository *Repository) GetByID(ctx context.Context, id string) (domainuser.User, error) {
	var userEntity domainuser.User
	query := persistence.LoadFile("repository/user/sql/get_by_id.sql")
	if err := repository.db.GetContext(ctx, &userEntity, query, id); err != nil {
		return domainuser.User{}, fmt.Errorf("get user %q: %w", id, err)
	}

	return userEntity, nil
}
