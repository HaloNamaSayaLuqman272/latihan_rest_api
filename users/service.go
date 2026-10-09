package users

import (
	"context"
	"latihan_rest_api/db/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (models.User, error)
	UpdateProfile(ctx context.Context, editReq *EditProfileRequest, id uuid.UUID) (models.User, error)
}

type service struct {
	get
	update
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB) Service {
	return service{
		get:    get{repository: repository},
		update: update{repository: repository, get: get{repository: repository}},
	}
}
