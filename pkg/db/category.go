package db

import (
	"github.com/adi221/good-reads/pkg/model"
)

type CategoryRespository interface {
	CreateCategoryForUser(uid uint, category model.Category) (*model.Category, error)
	GetCategoryByID(uid uint) (*model.Category, error)
	GetCategoriesByUser(uid uint, filter model.FilterSchema) (*model.GetCategoriesResponse, error)
	CountCategoriesByUser(uid uint) (uint, error)
}
