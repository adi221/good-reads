package db

import (
	"github.com/adi221/good-reads/pkg/model"
)

type ArticleRepository interface {
	CreateArticleForUser(uid uint, form model.ArticleCreateForm) (*model.Article, error)
	GetArticleByID(id uint) (*model.Article, error)
}
