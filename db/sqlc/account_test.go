package db

import (
	"context"
	"go-backend-course/simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(
		context.Background(), arg,
	)
	require.NoError(t, err, "Error")
	require.NotEmpty(t, account, "New row is empty")
	require.Equal(t, arg.Owner, account.Owner, "Account owner inserted is wrong")
	require.Equal(t, arg.Balance, account.Balance, "Account balance inserted is wrong")
	require.Equal(t, arg.Currency, account.Currency, "Account currency inserted is wrong")
	require.NotZero(t, account.ID, "Account id is zero")
	require.NotZero(t, account.CreatedAt, "Account createdAt is zero")

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	result, err := testQueries.GetAccountById(
		context.Background(), account.ID,
	)

	require.NoError(t, err)
	require.Equal(t, account.ID, result.ID)
}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	newBalance := util.RandomBalance()
	if account.Balance == newBalance {
		newBalance += 10
	}

	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: newBalance,
	}

	result, err := testQueries.UpdateAccount(
		context.Background(), args,
	)

	require.NoError(t, err)
	require.Equal(t, result, int64(1))

	updatedAccount, _ := testQueries.GetAccountById(
		context.Background(), account.ID,
	)
	require.Equal(t, newBalance, updatedAccount.Balance)
}
