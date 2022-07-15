// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateBlogArticle(ctx context.Context, arg CreateBlogArticleParams) (sql.Result, error)
	CreateBlogTag(ctx context.Context, arg CreateBlogTagParams) (sql.Result, error)
	DeleteArticle(ctx context.Context, arg DeleteArticleParams) (sql.Result, error)
	DeleteBlogTag(ctx context.Context, arg DeleteBlogTagParams) (sql.Result, error)
	GetBlogArticles(ctx context.Context, id int32) (BlogArticle, error)
	GetBlogTag(ctx context.Context, id int32) (BlogTag, error)
	ListBlogAtricles(ctx context.Context, arg ListBlogAtriclesParams) ([]BlogArticle, error)
	ListBlogTag(ctx context.Context, arg ListBlogTagParams) ([]BlogTag, error)
	UpdateBlogArticle(ctx context.Context, arg UpdateBlogArticleParams) (sql.Result, error)
	UpdateBlogTag(ctx context.Context, arg UpdateBlogTagParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
