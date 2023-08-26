package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/YuanData/SharedBoard/util"
	"github.com/stretchr/testify/require"
)

func createRandomSharedlink(t *testing.T) Sharedlink {
	arg := CreateSharedlinkParams{
		Name:    util.RandomSentence(),
		Urlhash: util.RandomUUID(),
	}

	sharedlink, err := testQueries.CreateSharedlink(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sharedlink)

	require.Equal(t, arg.Name, sharedlink.Name)
	require.Equal(t, arg.Urlhash, sharedlink.Urlhash)

	require.NotZero(t, sharedlink.ID)
	require.NotZero(t, sharedlink.CreatedAt)

	return sharedlink
}

func TestCreateSharedlink(t *testing.T) {
	createRandomSharedlink(t)
}

func TestGetSharedlink(t *testing.T) {
	sharedlink1 := createRandomSharedlink(t)
	sharedlink2, err := testQueries.GetSharedlink(context.Background(), sharedlink1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, sharedlink2)

	require.Equal(t, sharedlink1.ID, sharedlink2.ID)
	require.Equal(t, sharedlink1.Name, sharedlink2.Name)
	require.Equal(t, sharedlink1.Urlhash, sharedlink2.Urlhash)
	require.WithinDuration(t, sharedlink1.CreatedAt, sharedlink2.CreatedAt, time.Second)
}

func TestUpdateSharedlink(t *testing.T) {
	sharedlink1 := createRandomSharedlink(t)

	arg := UpdateSharedlinkParams{
		ID:   sharedlink1.ID,
		Name: util.RandomSentence(),
	}

	sharedlink2, err := testQueries.UpdateSharedlink(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sharedlink2)

	require.Equal(t, sharedlink1.ID, sharedlink2.ID)
	require.Equal(t, sharedlink1.Urlhash, sharedlink2.Urlhash)
	require.Equal(t, arg.Name, sharedlink2.Name)
	require.WithinDuration(t, sharedlink1.CreatedAt, sharedlink2.CreatedAt, time.Second)
}

func TestDeleteSharedlink(t *testing.T) {
	sharedlink1 := createRandomSharedlink(t)
	err := testQueries.DeleteSharedlink(context.Background(), sharedlink1.ID)
	require.NoError(t, err)

	sharedlink2, err := testQueries.GetSharedlink(context.Background(), sharedlink1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, sharedlink2)
}

func TestListSharedlinks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSharedlink(t)
	}

	arg := ListSharedlinkParams{
		Limit:  5,
		Offset: 0,
	}

	sharedlinks, err := testQueries.ListSharedlink(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sharedlinks)

	for _, sharedlink := range sharedlinks {
		require.NotEmpty(t, sharedlink)
	}
}
