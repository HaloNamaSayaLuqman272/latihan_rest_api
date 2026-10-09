package users

import (
	"context"
	"latihan_rest_api/db/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type get struct {
	repository *gorm.DB
}

func (g get) GetProfile(ctx context.Context, userID uuid.UUID) (models.User, error) {
	user := new(models.User)

	err := g.repository.WithContext(ctx).First(user, "id = ?", userID).Error
	if err != nil {
		return models.User{}, err
	}

	return *user, nil
}
