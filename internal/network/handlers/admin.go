package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/usecase"
	"github.com/v.kirpichov/admin/pkg/errorHandler"
)

type AdminHandler struct {
	cfg          *configs.Config
	adminUsecase usecase.AdminU
}

func NewAdminHandler(adminUsecase usecase.AdminU, cfg *configs.Config) *AdminHandler {
	return &AdminHandler{
		adminUsecase: adminUsecase,
		cfg:          cfg,
	}
}

func (a *AdminHandler) SignIn(c echo.Context) error {
	ctx := c.Request().Context()

	admin := &models.Admin{}
	err := json.NewDecoder(c.Request().Body).Decode(admin)
	if err != nil {
		c.Error(errorHandler.ErrInternal)
		return errorHandler.ErrInternal
	}

	err = a.adminUsecase.SignIn(ctx, admin)
	if err != nil {
		c.Error(errorHandler.NewInternalServerError(err.Error()))
		return errorHandler.ErrUnauthorized
	}

	sessionToken, err := a.adminUsecase.GetSessionToken(ctx, admin.Username)
	if err != nil {
		c.Error(errorHandler.NewInternalServerError(err.Error()))
		return errorHandler.ErrInternal
	}

	sessionCookie := &http.Cookie{
		Name:     "admin_session",
		Value:    sessionToken,
		Expires:  time.Now().Add(time.Duration(a.cfg.AdminSessionConfig.ExpiresAt * int(time.Hour.Nanoseconds()))),
		Secure:   true,
		HttpOnly: true,
	}

	c.SetCookie(sessionCookie)
	c.NoContent(http.StatusOK)
	return nil
}
