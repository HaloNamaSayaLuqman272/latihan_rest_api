package auth

import (
	"context"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/rabbitmq"

	"gorm.io/gorm"
)

type Service interface {
	RegisterUser(ctx context.Context, registerRequest *RegisterRequest) (models.User, error)
	LoginUser(ctx context.Context, loginRequest *LoginRequest) (models.User, error)
}

type service struct {
	register
	login
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB, mq rabbitmq.Publisher) Service {
	return service{
		register: register{repository: repository, publisher: mq},
		login:    login{repository: repository},
	}
}
