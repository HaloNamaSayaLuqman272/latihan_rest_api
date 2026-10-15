package products

import (
	"context"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/utils"

	"gorm.io/gorm"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAllProducts(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	products := []models.Product{}

	if err := g.repository.WithContext(ctx).Scopes(utils.Paginate(&products, &pagination, g.repository)).Find(&products).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = products

	return pagination, nil
}
