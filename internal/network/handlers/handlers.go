package handlers

import (
	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/usecase"
)

type Handlers struct {
	UserHandler      *UserHandler
	AdminHandler     *AdminHandler
	ServicesHandler  *ServicesHandler
	EmployersHandler *EmployersHandler
	ReportsHandler   *ReportsHandler
}

func NewHandlers(usecases *usecase.Usecases, cfg *configs.Config) *Handlers {
	return &Handlers{
		UserHandler:      NewUserHandler(usecases.UserUsecase, cfg),
		AdminHandler:     NewAdminHandler(usecases.AdminUsecase, cfg),
		ServicesHandler:  NewServicesUsecase(usecases.ServiceUsecase),
		EmployersHandler: NewEmployersHandler(usecases.EmployersUsecases),
		ReportsHandler:   NewReportsHandler(usecases.ReportsUsecase),
	}
}
