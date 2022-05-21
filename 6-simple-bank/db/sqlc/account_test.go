package db

import (
	"context"
	"database/sql"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/require"
	"simple-bank/utils"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestQueries_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestQueries_GetAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)
	retrievedAccount, err := testQueries.GetAccount(context.Background(), randomAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	require.Equal(t, randomAccount.ID, retrievedAccount.ID)
	require.Equal(t, randomAccount.Owner, retrievedAccount.Owner)
	require.Equal(t, randomAccount.Balance, retrievedAccount.Balance)
	require.Equal(t, randomAccount.Currency, retrievedAccount.Currency)
	require.WithinDuration(t, randomAccount.CreatedAt, retrievedAccount.CreatedAt, time.Second)
}

func TestQueries_UpdateAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      randomAccount.ID,
		Balance: 500.45,
	}

	retrievedAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	require.Equal(t, randomAccount.ID, retrievedAccount.ID)
	require.Equal(t, randomAccount.Owner, retrievedAccount.Owner)
	require.Equal(t, arg.Balance, retrievedAccount.Balance)
	require.Equal(t, randomAccount.Currency, retrievedAccount.Currency)
	require.WithinDuration(t, randomAccount.CreatedAt, retrievedAccount.CreatedAt, time.Second)
}

func TestQueries_DeleteAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), randomAccount.ID)
	require.NoError(t, err)

	retrievedAccount, err := testQueries.GetAccount(context.Background(), randomAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, retrievedAccount)
}

func TestQueries_ListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
