// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createBlogArticleStmt, err = db.PrepareContext(ctx, createBlogArticle); err != nil {
		return nil, fmt.Errorf("error preparing query CreateBlogArticle: %w", err)
	}
	if q.createBlogTagStmt, err = db.PrepareContext(ctx, createBlogTag); err != nil {
		return nil, fmt.Errorf("error preparing query CreateBlogTag: %w", err)
	}
	if q.deleteArticleStmt, err = db.PrepareContext(ctx, deleteArticle); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteArticle: %w", err)
	}
	if q.deleteBlogTagStmt, err = db.PrepareContext(ctx, deleteBlogTag); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteBlogTag: %w", err)
	}
	if q.getBlogArticlesStmt, err = db.PrepareContext(ctx, getBlogArticles); err != nil {
		return nil, fmt.Errorf("error preparing query GetBlogArticles: %w", err)
	}
	if q.getBlogTagStmt, err = db.PrepareContext(ctx, getBlogTag); err != nil {
		return nil, fmt.Errorf("error preparing query GetBlogTag: %w", err)
	}
	if q.listBlogAtriclesStmt, err = db.PrepareContext(ctx, listBlogAtricles); err != nil {
		return nil, fmt.Errorf("error preparing query ListBlogAtricles: %w", err)
	}
	if q.listBlogTagStmt, err = db.PrepareContext(ctx, listBlogTag); err != nil {
		return nil, fmt.Errorf("error preparing query ListBlogTag: %w", err)
	}
	if q.updateBlogArticleStmt, err = db.PrepareContext(ctx, updateBlogArticle); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateBlogArticle: %w", err)
	}
	if q.updateBlogTagStmt, err = db.PrepareContext(ctx, updateBlogTag); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateBlogTag: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createBlogArticleStmt != nil {
		if cerr := q.createBlogArticleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createBlogArticleStmt: %w", cerr)
		}
	}
	if q.createBlogTagStmt != nil {
		if cerr := q.createBlogTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createBlogTagStmt: %w", cerr)
		}
	}
	if q.deleteArticleStmt != nil {
		if cerr := q.deleteArticleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteArticleStmt: %w", cerr)
		}
	}
	if q.deleteBlogTagStmt != nil {
		if cerr := q.deleteBlogTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteBlogTagStmt: %w", cerr)
		}
	}
	if q.getBlogArticlesStmt != nil {
		if cerr := q.getBlogArticlesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getBlogArticlesStmt: %w", cerr)
		}
	}
	if q.getBlogTagStmt != nil {
		if cerr := q.getBlogTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getBlogTagStmt: %w", cerr)
		}
	}
	if q.listBlogAtriclesStmt != nil {
		if cerr := q.listBlogAtriclesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listBlogAtriclesStmt: %w", cerr)
		}
	}
	if q.listBlogTagStmt != nil {
		if cerr := q.listBlogTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listBlogTagStmt: %w", cerr)
		}
	}
	if q.updateBlogArticleStmt != nil {
		if cerr := q.updateBlogArticleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateBlogArticleStmt: %w", cerr)
		}
	}
	if q.updateBlogTagStmt != nil {
		if cerr := q.updateBlogTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateBlogTagStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                    DBTX
	tx                    *sql.Tx
	createBlogArticleStmt *sql.Stmt
	createBlogTagStmt     *sql.Stmt
	deleteArticleStmt     *sql.Stmt
	deleteBlogTagStmt     *sql.Stmt
	getBlogArticlesStmt   *sql.Stmt
	getBlogTagStmt        *sql.Stmt
	listBlogAtriclesStmt  *sql.Stmt
	listBlogTagStmt       *sql.Stmt
	updateBlogArticleStmt *sql.Stmt
	updateBlogTagStmt     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                    tx,
		tx:                    tx,
		createBlogArticleStmt: q.createBlogArticleStmt,
		createBlogTagStmt:     q.createBlogTagStmt,
		deleteArticleStmt:     q.deleteArticleStmt,
		deleteBlogTagStmt:     q.deleteBlogTagStmt,
		getBlogArticlesStmt:   q.getBlogArticlesStmt,
		getBlogTagStmt:        q.getBlogTagStmt,
		listBlogAtriclesStmt:  q.listBlogAtriclesStmt,
		listBlogTagStmt:       q.listBlogTagStmt,
		updateBlogArticleStmt: q.updateBlogArticleStmt,
		updateBlogTagStmt:     q.updateBlogTagStmt,
	}
}
