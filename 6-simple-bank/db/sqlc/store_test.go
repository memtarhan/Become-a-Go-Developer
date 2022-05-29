package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"simple-bank/utils"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	fmt.Println("Testing Transfer...")
	store := NewStore(testDB)

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	fmt.Println("Before started: From account:", fromAccount.ID, "with balance:", fromAccount.Balance)
	fmt.Println("Before started: To account:", toAccount.ID, "with balance:", toAccount.Balance)

	// run n concurrent transfer transactions
	n := 2
	amount := 100.46

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountId: fromAccount.ID,
				ToAccountId:   toAccount.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromAccount.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toAccount.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check account
		transferFromAccount := result.FromAccount
		require.NotEmpty(t, transferFromAccount)
		require.Equal(t, transferFromAccount.ID, fromAccount.ID)

		transferToAccount := result.ToAccount
		require.NotEmpty(t, transferToAccount)
		require.Equal(t, transferToAccount.ID, toAccount.ID)

		// check accounts' balance
		difference1 := fromAccount.Balance - transferFromAccount.Balance
		difference2 := transferToAccount.Balance - toAccount.Balance
		require.Equal(t, utils.RoundFloat(difference1), utils.RoundFloat(difference2))
		require.True(t, difference1 > 0)
		//require.True(t, difference1%amount == 0)

		k := int(difference1 / amount)
		//require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balances
	updatedFromAccount, err := testQueries.GetAccount(context.Background(), fromAccount.ID)
	require.NoError(t, err)

	updatedToAccount, err := testQueries.GetAccount(context.Background(), toAccount.ID)
	require.NoError(t, err)

	count := float64(n)
	require.Equal(t, utils.RoundFloat(fromAccount.Balance-(amount*count)), utils.RoundFloat(updatedFromAccount.Balance))
	require.Equal(t, toAccount.Balance+(amount*count), utils.RoundFloat(updatedToAccount.Balance))

	fmt.Println("After finished: From account:", updatedFromAccount.ID, "with balance:", updatedFromAccount.Balance)
	fmt.Println("After finished: To account:", updatedToAccount.ID, "with balance:", updatedToAccount.Balance)
}
