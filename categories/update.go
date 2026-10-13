package categories

import (
	"context"
	"latihan_rest_api/db/models"

	"gorm.io/gorm"
)

type update struct {
	repository *gorm.DB
	getbyid
}

func (u update) UpdateCategory(ctx context.Context, categoryReq *CategoryRequest, id uint) (models.Category, error) {
	category := models.Category{
		CategoryName: categoryReq.CategoryName,
		Description:  categoryReq.Description,
	}

	result := u.repository.WithContext(ctx).Where("id = ?", id).Updates(&category)
	if err := result.Error; err != nil {
		return models.Category{}, err
	}

	record, err := u.getbyid.GetCategoryByID(ctx, id)
	if err != nil {
		return models.Category{}, err
	}

	return record, nil
}
