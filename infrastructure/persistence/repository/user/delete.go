package user

import (
	"context"
	"fmt"

	domainuser "Momentum/domain/user"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Delete(ctx context.Context, id string) (domainuser.User, error) {
	var userEntity domainuser.User
	query := persistence.LoadFile("repository/user/sql/delete.sql")
	if err := repository.db.GetContext(ctx, &userEntity, query, id); err != nil {
		return domainuser.User{}, fmt.Errorf("delete user %q: %w", id, err)
	}

	return userEntity, nil
}
