package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v; rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	Sender   int64 `json:"sender"`
	Receiver int64 `json:"receiver"`
	Amount   int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		transfer, err := q.CreateTransfer(ctx, CreateTransferParams{
			Sender:   arg.Sender,
			Receiver: arg.Receiver,
			Amount:   arg.Amount,
		})

		if err != nil {
			return err
		}

		result.Transfer = transfer

		senderUpdated, _ := q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.Sender,
			Amount: -arg.Amount,
		})

		if senderUpdated != 1 {
			return fmt.Errorf("sender account %d not found", arg.Sender)
		}

		receiverUpdated, err := q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.Receiver,
			Amount: arg.Amount,
		})

		if receiverUpdated != 1 {
			return fmt.Errorf("receiver account %d not found", arg.Receiver)
		}

		return nil
	})

	return result, err
}
