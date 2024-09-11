package db

import (
	"context"
	"github.com/SaishNaik/simplebank/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {

	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	params := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}
	ctx := context.Background()
	user, err := testQueries.CreateUser(ctx, params)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, params.Username, user.Username)
	require.Equal(t, params.HashedPassword, user.HashedPassword)
	require.Equal(t, params.Username, user.Username)
	require.Equal(t, params.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	createdUser := createRandomUser(t)
	gotUser, err := testQueries.GetUser(ctx, createdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, gotUser)
	require.Equal(t, createdUser.Username, gotUser.Username)
	require.Equal(t, createdUser.HashedPassword, gotUser.HashedPassword)
	require.Equal(t, createdUser.Username, gotUser.Username)
	require.Equal(t, createdUser.Email, gotUser.Email)
	require.WithinDuration(t, createdUser.CreatedAt, gotUser.CreatedAt, time.Second)
	require.WithinDuration(t, createdUser.PasswordChangedAt, gotUser.PasswordChangedAt, time.Second)
}
