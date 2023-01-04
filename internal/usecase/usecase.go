package usecase

import (
	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/repository"
)

type Usecases struct {
	UserUsecase UserUsecase
}

func NewUsecases(repository *repository.Repository, cfg *configs.Config) *Usecases {
	return &Usecases{
		UserUsecase: *NewUserUsecase(&repository.Users, cfg),
	}
}
