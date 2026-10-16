package products

import (
	"context"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type getbycategory struct {
	repository *gorm.DB
}

func (g getbycategory) GetProductsByCategory(ctx context.Context, pagination utils.Pagination, categoryID uint) (utils.Pagination, error) {
	products := []models.Product{}

	if err := g.repository.WithContext(ctx).Scopes(utils.PaginateByProductCategory(&products, &pagination, categoryID, g.repository)).Preload(clause.Associations).Find(&products).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = products

	return pagination, nil
}
