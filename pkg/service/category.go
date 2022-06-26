package service

import (
	"context"

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
