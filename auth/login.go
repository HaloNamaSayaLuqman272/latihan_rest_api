package auth

import (
	"context"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/utils"

	"gorm.io/gorm"
)

type login struct {
	repository *gorm.DB
}

func (l login) LoginUser(ctx context.Context, loginRequest *LoginRequest) (models.User, error) {
	user := new(models.User)
	if err := l.repository.WithContext(ctx).First(user, "email = ?", loginRequest.Email).Error; err != nil {
		// ini adalah proses mencocokkan email login user
		return models.User{}, err
	}

	if err := utils.ComparePassword(user.Password, loginRequest.Password); err != nil {
		// ini adalah proses membandingkan kecocokan "password register" dan "password login"
		return models.User{}, err
	}

	return *user, nil
}
