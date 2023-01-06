package handlers

import (
	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/usecase"
)

type Handlers struct {
	UserHandler     *UserHandler
	AdminHandler    *AdminHandler
	ServicesHandler *ServicesHandler
}

func NewHandlers(usecases *usecase.Usecases, cfg *configs.Config) *Handlers {
	return &Handlers{
		UserHandler:     NewUserHandler(usecases.UserUsecase, cfg),
		AdminHandler:    NewAdminHandler(usecases.AdminUsecase, cfg),
		ServicesHandler: NewServicesUsecase(usecases.ServiceUsecase),
	}
}
