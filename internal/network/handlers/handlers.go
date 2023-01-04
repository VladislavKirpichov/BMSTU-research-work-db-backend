package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/usecase"
)

type Handlers struct {
	UserHandler UserH
}

func NewHandlers(usecases *usecase.Usecases, cfg *configs.Config) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(&usecases.UserUsecase, cfg),
	}
}

// Write new interfaces here
type UserH interface {
	SignIn(c echo.Context) error
	SignUp(c echo.Context) error
	GetAllUsers(c echo.Context) error
}
