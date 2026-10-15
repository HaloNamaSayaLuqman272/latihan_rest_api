package products

import (
	"context"
	"latihan_rest_api/db/models"

	"gorm.io/gorm"
)

type getbyid struct {
	repository *gorm.DB
}

func (g getbyid) GetProductByID(ctx context.Context, id uint) (models.Product, error) {
	product := new(models.Product)
	if err := g.repository.WithContext(ctx).First(product, "id = ?", id).Error; err != nil {
		return models.Product{}, err
	}

	return *product, nil
}
