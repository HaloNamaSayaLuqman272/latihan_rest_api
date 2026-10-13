package categories

import (
	"context"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/utils"

	"gorm.io/gorm"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAllCategories(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	categories := []models.Category{}

	if err := g.repository.WithContext(ctx).Scopes(utils.Paginate(&categories, &pagination, g.repository)).Find(&categories).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = categories

	return pagination, nil
}
