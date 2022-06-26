package db

import (
	"github.com/adi221/good-reads/pkg/model"
)

type CategoryRespository interface {
	CreateCategoryForUser(uid uint, category model.Category) (*model.Category, error)
}
