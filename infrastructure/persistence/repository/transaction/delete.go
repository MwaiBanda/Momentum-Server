package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domaintransaction "Momentum/domain/transaction"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Delete(ctx context.Context, id string) (domaintransaction.Transaction, error) {
	var row transactionRow
	query := persistence.LoadFile("repository/transaction/sql/delete.sql")
	if err := repository.db.GetContext(ctx, &row, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domaintransaction.Transaction{}, fmt.Errorf("delete transaction %q: %w", id, domaintransaction.ErrNotFound)
		}
		return domaintransaction.Transaction{}, fmt.Errorf("delete transaction %q: %w", id, err)
	}
	return row.entity(), nil
}
