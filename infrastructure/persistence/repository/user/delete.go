package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domainuser "Momentum/domain/user"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Delete(ctx context.Context, id string) (domainuser.User, error) {
	var userEntity domainuser.User
	query := persistence.LoadFile("repository/user/sql/delete.sql")
	if err := repository.db.GetContext(ctx, &userEntity, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainuser.User{}, fmt.Errorf("delete user %q: %w", id, domainuser.ErrNotFound)
		}
		return domainuser.User{}, fmt.Errorf("delete user %q: %w", id, err)
	}

	return userEntity, nil
}
