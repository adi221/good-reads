package postgres

import (
	"database/sql"
	"strings"

	"github.com/adi221/good-reads/pkg/model"
)

const articlesTable = "articles"

var articleColumns = []string{
	"id",
	"\"userId\"",
	"\"categoryId\"",
	"title",
	"text",
	"html",
	"url",
	"image",
	"hash",
	"status",
	"\"createdAt\"",
	"\"updatedAt\"",
}

func mapRowToArticle(row *sql.Row) (*model.Article, error) {
	article := &model.Article{}

	err := row.Scan(
		&article.ID,
		&article.UserID,
		&article.CategoryID,
		&article.Title,
		&article.Text,
		&article.HTML,
		&article.URL,
		&article.Image,
		&article.Hash,
		&article.Status,
		&article.CreatedAt,
		&article.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, mapError(err)
	}
	return article, nil
}

// CreateArticleForUser creates an article into the DB
func (pg *DB) CreateArticleForUser(uid uint, form model.ArticleCreateForm) (*model.Article, error) {
	query, args, _ := pg.psql.Insert(
		articlesTable,
	).Columns(
		"userId",
		"categoryId",
		"title",
		"text",
		"html",
		"url",
		"image",
		"hash",
		"status",
		"updatedAt",
	).Values(
		uid,
		form.CategoryID,
		form.Title,
		form.Text,
		form.HTML,
		form.URL,
		form.Image,
		form.Hash(),
		"unread",
		"NOW()",
	).Suffix(
		"RETURNING " + strings.Join(articleColumns, ","),
	).ToSql()
	row := pg.db.QueryRow(query, args...)
	return mapRowToArticle(row)
}
