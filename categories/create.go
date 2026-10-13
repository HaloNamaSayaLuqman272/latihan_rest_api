package categories

import (
	"context"
	"latihan_rest_api/db/models"

	"gorm.io/gorm"
)

type create struct {
	repository *gorm.DB
}

func (c create) CreateCategory(ctx context.Context, categoryReq *CategoryRequest) (models.Category, error) {
	category := models.Category{
		CategoryName: categoryReq.CategoryName,
		Description:  categoryReq.Description,
	}

	result := c.repository.WithContext(ctx).Create(&category)
	record := new(models.Category)

	if err := result.Error; err != nil {
		return models.Category{}, err
	}

	if err := result.WithContext(ctx).Last(record).Error; err != nil {
		return models.Category{}, err
	}

	return *record, nil
}
