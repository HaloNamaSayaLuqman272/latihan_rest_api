package categories

import (
	"context"

	"gorm.io/gorm"
)

type delete struct {
	repository *gorm.DB
	getbyid
}

func (d delete) DeleteCategory(ctx context.Context, id uint) error {
	category, err := d.getbyid.GetCategoryByID(ctx, id)
	if err != nil {
		return err
	}

	if err := d.repository.WithContext(ctx).Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
