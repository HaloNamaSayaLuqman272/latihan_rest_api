package auth

import (
	"context"

	"gorm.io/gorm"
)

type Service interface {
	RegisterUser(ctx context.Context, registerRequest *RegisterRequest) (User, error)
	LoginUser(ctx context.Context, loginRequest *LoginRequest) (User, error)
}

type service struct {
	register
	login
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB) Service {
	return service{
		register: register{repository: repository},
		login:    login{repository: repository},
	}
}
