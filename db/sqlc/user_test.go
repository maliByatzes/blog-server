package db

import (
	"context"
	"testing"
	"time"

	"github.com/maliByatzes/blog-server/util"
	"github.com/stretchr/testify/require"
)

func createRandomuser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(12))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.UpdatedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	require.NotEmpty(t, user.RoleID)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomuser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomuser(t)
	require.NotEmpty(t, user1)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.RoleID, user2.RoleID)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomuser(t)
	require.NotEmpty(t, user1)

	hashedPassword, err := util.HashPassword(util.RandomString(12))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := UpdateUserParams{
		Username:       user1.Username,
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.HashedPassword, user2.HashedPassword)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, user1.RoleID, user2.RoleID)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
