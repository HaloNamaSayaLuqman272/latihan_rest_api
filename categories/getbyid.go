package categories

import (
	"context"
	"latihan_rest_api/db/models"

	"gorm.io/gorm"
)

type getbyid struct {
	repository *gorm.DB
}

func (g getbyid) GetCategoryByID(ctx context.Context, id uint) (models.Category, error) {
	category := new(models.Category)
	if err := g.repository.WithContext(ctx).First(category, "id = ?", id).Error; err != nil {
		return models.Category{}, err
	}

	return *category, nil
}
