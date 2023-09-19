package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/maliByatzes/blog-server/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		CategoryName: util.RandomCategory(),
		Description:  util.RandomString(50),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.CategoryName, category.CategoryName)
	require.Equal(t, arg.Description, category.Description)
	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	require.NotEmpty(t, category1)

	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.CategoryName, category2.CategoryName)
	require.Equal(t, category1.Description, category2.Description)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestListCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCategory(t)
	}

	categories, err := testQueries.ListCategories(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	require.NotEmpty(t, category1)

	arg := UpdateCategoryParams{
		ID:           category1.ID,
		CategoryName: util.RandomCategory(),
		Description:  util.RandomString(50),
	}

	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.CategoryName, category2.CategoryName)
	require.Equal(t, arg.Description, category2.Description)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	require.NotEmpty(t, category1)

	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}
