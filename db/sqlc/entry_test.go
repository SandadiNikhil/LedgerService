package db

import (
	"context"
	"testing"

	"github.com/SandadiNikhil/LedgerService/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, fromAccount, toAccount Account) Entry {
	arg := CreateEntryParams{
		AccountID: fromAccount.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	createRandomEntry(t, fromAccount, toAccount)
}

func TestGetEntry(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	entry1 := createRandomEntry(t, fromAccount, toAccount)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, 0)
}

func TestListEntries(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, fromAccount, toAccount)
	}

	arg := ListEntriesParams{
		AccountID: fromAccount.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, fromAccount.ID, entry.AccountID)
	}
}
