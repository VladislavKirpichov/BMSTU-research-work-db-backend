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

func (a *AdminHandler) Auth(c echo.Context) error {
	cookies := c.Cookies()
	ctx := c.Request().Context()

	for _, cookie := range cookies {
		if cookie.Name == "admin_session" {
			if _, err := a.adminUsecase.Auth(ctx, cookie.Value); err != nil {
				return errorHandler.ErrInvalidSession
			} else {
				return nil
			}
		}
	}

	return errorHandler.ErrInvalidSession
}

type AppliesResponse struct {
	Applies []*models.ApplyWithData `json:"applies,omitempty"`
}

func (a *AdminHandler) Applies(c echo.Context) error {
	ctx := c.Request().Context()

	applies, err := a.adminUsecase.GetAllApplies(ctx)
	if err != nil {
		return errorHandler.ErrInternal
	}

	c.JSON(http.StatusOK, &AppliesResponse{
		Applies: applies,
	})

	return nil
}

func (a *AdminHandler) Logout(c echo.Context) error {
	ctx := c.Request().Context()
	cookies := c.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "admin_session" {
			err := a.adminUsecase.Logout(ctx, cookie.Value)
			if err != nil {
				return errorHandler.ErrInternal
			}

			return nil
		}
	}

	return errorHandler.ErrInvalidSession
}
