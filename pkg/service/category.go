package service

import (
	"context"
	"net/http"

	"github.com/adi221/good-reads/pkg/helper"
	"github.com/adi221/good-reads/pkg/model"
)

func (reg *Registry) CreateCategory(ctx context.Context, category model.Category) (*model.Category, error) {
	uid := getCurrentUserIDFromContext(ctx)
	result, err := reg.db.CreateCategoryForUser(uid, category)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (reg *Registry) GetCategory(ctx context.Context, id uint) (*model.Category, error) {
	uid := getCurrentUserIDFromContext(ctx)
	category, err := reg.db.GetCategoryByID(id)
	if err != nil || category == nil || *category.UserID != uid {
		return nil, helper.NewHttpError(
			model.CategoryNonExist,
			http.StatusNotFound,
			"Category not found",
		)
	}
	return category, nil
}
