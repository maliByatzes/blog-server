package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/maliByatzes/blog-server/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) Post {
	user := createRandomuser(t)
	require.NotEmpty(t, user)

	category := createRandomCategory(t)
	require.NotEmpty(t, category)

	arg := CreatePostParams{
		Username:   user.Username,
		Title:      util.RandomString(12),
		ImageUrl:   util.RandomString(20),
		Content:    util.RandomString(255),
		CategoryID: category.ID,
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.Username, post.Username)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.ImageUrl, post.ImageUrl)
	require.Equal(t, arg.Content, post.Content)
	require.Equal(t, arg.CategoryID, post.CategoryID)
	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)

	return post
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPost(t *testing.T) {
	post1 := createRandomPost(t)
	require.NotEmpty(t, post1)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.Username, post2.Username)
	require.Equal(t, post2.Title, post2.Title)
	require.Equal(t, post1.ImageUrl, post2.ImageUrl)
	require.Equal(t, post1.Content, post2.Content)
	require.Equal(t, post1.CategoryID, post2.CategoryID)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestListPosts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPost(t)
	}

	arg := ListPostsParams{
		Limit:  5,
		Offset: 5,
	}

	posts, err := testQueries.ListPosts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}

func TestUpdatePost(t *testing.T) {
	post1 := createRandomPost(t)
	require.NotEmpty(t, post1)

	arg := UpdatePostParams{
		ID:         post1.ID,
		Title:      util.RandomString(12),
		ImageUrl:   util.RandomString(20),
		Content:    util.RandomString(255),
		CategoryID: post1.CategoryID,
	}

	post2, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.Username, post2.Username)
	require.Equal(t, arg.Title, post2.Title)
	require.Equal(t, arg.ImageUrl, post2.ImageUrl)
	require.Equal(t, arg.Content, post2.Content)
	require.Equal(t, post1.CategoryID, post2.CategoryID)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestDeletePost(t *testing.T) {
	post1 := createRandomPost(t)
	require.NotEmpty(t, post1)

	err := testQueries.DeletePost(context.Background(), post1.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}
