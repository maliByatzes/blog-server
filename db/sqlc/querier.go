// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateRole(ctx context.Context, roleName string) (Role, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCategory(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	GetCategory(ctx context.Context, id int64) (Category, error)
	GetPost(ctx context.Context, id int64) (Post, error)
	GetRole(ctx context.Context, id int64) (Role, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListCategories(ctx context.Context) ([]Category, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
