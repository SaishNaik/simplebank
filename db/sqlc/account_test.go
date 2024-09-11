package db

import (
	"context"
	"database/sql"
	"github.com/SaishNaik/simplebank/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	params := CreateAccountParams{
		Owner:    user.Username,
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	ctx := context.Background()
	account, err := testQueries.CreateAccount(ctx, params)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, params.Owner, account.Owner)
	require.Equal(t, params.Balance, account.Balance)
	require.Equal(t, params.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	ctx := context.Background()
	createdAccount := createRandomAccount(t)
	gotAccount, err := testQueries.GetAccount(ctx, createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)
	require.Equal(t, createdAccount.Owner, gotAccount.Owner)
	require.Equal(t, createdAccount.Balance, gotAccount.Balance)
	require.Equal(t, createdAccount.Currency, gotAccount.Currency)
	require.Equal(t, createdAccount.ID, gotAccount.ID)
	require.WithinDuration(t, createdAccount.CreatedAt, gotAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	expectedBalance := utils.RandomMoney()
	params := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: expectedBalance,
	}
	gotAccount, err := testQueries.UpdateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)

	require.Equal(t, createdAccount.Owner, gotAccount.Owner)
	require.Equal(t, createdAccount.Currency, gotAccount.Currency)
	require.Equal(t, createdAccount.ID, gotAccount.ID)
	require.WithinDuration(t, createdAccount.CreatedAt, gotAccount.CreatedAt, time.Second)
	require.Equal(t, expectedBalance, gotAccount.Balance)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	gotAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, gotAccount)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	params := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, len(accounts), 5)
	for i := 0; i < len(accounts); i++ {
		require.NotEmpty(t, accounts[i])
	}
}
