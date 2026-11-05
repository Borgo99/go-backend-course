package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleTransfer(t *testing.T) {
	store := NewStore(testDB)
	q := store.Queries

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	const newBalanceForAccount1 = int64(200)
	const transferAmount = int64(100)

	_, err := q.UpdateAccount(
		context.Background(), UpdateAccountParams{
			ID:      account1.ID,
			Balance: newBalanceForAccount1,
		},
	)
	require.NoError(t, err)
	account1.Balance = newBalanceForAccount1

	transferResult, err := store.TransferTx(
		context.Background(), TransferTxParams{
			Sender:   account1.ID,
			Receiver: account2.ID,
			Amount:   transferAmount,
		},
	)
	require.NoError(t, err)
	require.Equal(t, account1.ID, transferResult.Transfer.Sender)

	account1Updated, _ := q.GetAccountById(
		context.Background(), account1.ID,
	)
	require.Equal(t, newBalanceForAccount1-transferAmount, account1Updated.Balance)

	account2Updated, _ := q.GetAccountById(
		context.Background(), account2.ID,
	)
	require.Equal(t, account2.Balance+transferAmount, account2Updated.Balance)

}

func TestConcurrentTransfers(t *testing.T) {
	store := NewStore(testDB)
	q := store.Queries

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	const newBalanceForAccount1 = int64(200)
	const transferAmount = int64(10)

	_, err := q.UpdateAccount(
		context.Background(), UpdateAccountParams{
			ID:      account1.ID,
			Balance: newBalanceForAccount1,
		},
	)
	require.NoError(t, err)
	account1.Balance = newBalanceForAccount1

	errs := make(chan error)
	results := make(chan TransferTxResult)
	n := 10

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(
				context.Background(), TransferTxParams{
					Sender:   account1.ID,
					Receiver: account2.ID,
					Amount:   transferAmount,
				},
			)
			results <- result
			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		res := <-results
		require.NotEmpty(t, res)
		err := <-errs
		require.NoError(t, err)
	}

	account1Updated, _ := q.GetAccountById(
		context.Background(), account1.ID,
	)
	require.Equal(t, newBalanceForAccount1-(transferAmount*int64(n)), account1Updated.Balance)

	account2Updated, _ := q.GetAccountById(
		context.Background(), account2.ID,
	)
	require.Equal(t, account2.Balance+(transferAmount*int64(n)), account2Updated.Balance)

}
