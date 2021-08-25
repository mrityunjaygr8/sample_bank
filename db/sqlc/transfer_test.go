package db

import (
	"context"
	"testing"

	"github.com/mrityunjaygr8/sample_bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, from_account, to_account Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from_account.ID,
		ToAccountID:   to_account.ID,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, from_account.ID, transfer.FromAccountID)
	require.Equal(t, to_account.ID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)

	createRandomTransfer(t, from_account, to_account)
}

func TestGetTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, from_account, to_account)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.CreatedAt, transfer2.CreatedAt)

}

func TestListTransfers(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, from_account, to_account)
		createRandomTransfer(t, to_account, from_account)
	}

	args := ListTransfersParams{
		FromAccountID: from_account.ID,
		ToAccountID:   from_account.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}
