package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions tro execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates and retrieves a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function withing a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountId int64   `json:"from_account_id"`
	ToAccountId   int64   `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct {
}{}

// TransferTx performs a money transfer from one account to the other
// It creates a transfer record, add account entries and update accounts' balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		fmt.Println(txName, "create transfer")
		// Creating a transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "create from entry")
		// Creating from entry
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "create to entry")
		// Creating to entry
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "get from account")
		fromAccount, err := q.GetAccountForUpdate(ctx, arg.FromAccountId)
		if err != nil {
			return err
		}
		fmt.Println(txName, "update from account")
		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountId,
			Balance: fromAccount.Balance - arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "get to account")
		toAccount, err := q.GetAccountForUpdate(ctx, arg.ToAccountId)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update to account")
		result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.ToAccountId,
			Balance: toAccount.Balance + arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
