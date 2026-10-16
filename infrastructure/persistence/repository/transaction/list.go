package transaction

import (
	"context"
	"fmt"

	domaintransaction "Momentum/domain/transaction"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) List(ctx context.Context) ([]domaintransaction.Transaction, error) {
	var rows []transactionRow
	query := persistence.LoadFile("repository/transaction/sql/list.sql")
	if err := repository.db.SelectContext(ctx, &rows, query); err != nil {
		return nil, fmt.Errorf("list transactions: %w", err)
	}
	return mapRows(rows), nil
}

func (repository *Repository) ListByUser(ctx context.Context, userID string) ([]domaintransaction.Transaction, error) {
	var rows []transactionRow
	query := persistence.LoadFile("repository/transaction/sql/list_by_user.sql")
	if err := repository.db.SelectContext(ctx, &rows, query, userID); err != nil {
		return nil, fmt.Errorf("list transactions for user %q: %w", userID, err)
	}
	return mapRows(rows), nil
}

func mapRows(rows []transactionRow) []domaintransaction.Transaction {
	transactions := make([]domaintransaction.Transaction, 0, len(rows))
	for _, row := range rows {
		transactions = append(transactions, row.entity())
	}
	return transactions
}
