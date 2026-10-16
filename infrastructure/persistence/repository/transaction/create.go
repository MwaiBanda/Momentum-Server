package transaction

import (
	"context"
	"fmt"

	domaintransaction "Momentum/domain/transaction"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Create(ctx context.Context, params domaintransaction.CreateParams) (domaintransaction.Transaction, error) {
	transaction, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		return domaintransaction.Transaction{}, fmt.Errorf("begin transaction creation: %w", err)
	}
	defer transaction.Rollback()

	createUserQuery := persistence.LoadFile("repository/transaction/sql/create_user_if_missing.sql")
	if _, err := transaction.ExecContext(
		ctx,
		createUserQuery,
		params.UserID,
		params.Email,
		params.Fullname,
		params.Phone,
	); err != nil {
		return domaintransaction.Transaction{}, fmt.Errorf("create transaction user %q: %w", params.UserID, err)
	}

	createTransactionQuery := persistence.LoadFile("repository/transaction/sql/create.sql")
	if _, err := transaction.ExecContext(
		ctx,
		createTransactionQuery,
		params.ID,
		params.Amount,
		params.Date,
		params.UserID,
		params.Description,
	); err != nil {
		return domaintransaction.Transaction{}, fmt.Errorf("create transaction %q: %w", params.ID, err)
	}

	var row transactionRow
	getQuery := persistence.LoadFile("repository/transaction/sql/get_by_id.sql")
	if err := transaction.GetContext(ctx, &row, getQuery, params.ID); err != nil {
		return domaintransaction.Transaction{}, fmt.Errorf("get created transaction %q: %w", params.ID, err)
	}

	if err := transaction.Commit(); err != nil {
		return domaintransaction.Transaction{}, fmt.Errorf("commit transaction %q: %w", params.ID, err)
	}
	return row.entity(), nil
}
