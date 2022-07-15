// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"time"
)

type BlogArticle struct {
	ID int32 `json:"id"`
	// 标签ID
	TagID sql.NullInt32 `json:"tag_id"`
	// 文章标题
	Title sql.NullString `json:"title"`
	// 简述
	Desc    sql.NullString `json:"desc"`
	Content sql.NullString `json:"content"`
	// 创建时间
	CreatedOn time.Time `json:"created_on"`
	// 创建人
	CreatedBy sql.NullString `json:"created_by"`
	// 修改时间
	ModifiedOn time.Time `json:"modified_on"`
	// 修改人
	ModifiedBy sql.NullString `json:"modified_by"`
	DeletedOn  time.Time      `json:"deleted_on"`
	// 状态 0为禁用1为启用
	State sql.NullInt32 `json:"state"`
}

type BlogAuth struct {
	ID int32 `json:"id"`
	// 账号
	Username sql.NullString `json:"username"`
	// 密码
	Password sql.NullString `json:"password"`
	// 创建时间
	CreatedOn time.Time `json:"created_on"`
}

type BlogTag struct {
	ID int32 `json:"id"`
	// 标签名称
	Name sql.NullString `json:"name"`
	// 创建时间
	CreatedOn time.Time `json:"created_on"`
	// 创建人
	CreatedBy sql.NullString `json:"created_by"`
	// 修改时间
	ModifiedOn time.Time `json:"modified_on"`
	// 修改人
	ModifiedBy sql.NullString `json:"modified_by"`
	DeletedOn  time.Time      `json:"deleted_on"`
	// 状态 0为禁用、1为启用
	State sql.NullInt32 `json:"state"`
}
