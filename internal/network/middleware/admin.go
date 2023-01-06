package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/internal/repository"
	"github.com/v.kirpichov/admin/pkg/errorHandler"
)

type AdminSession struct {
	sessionRepo repository.SessionR
}

func NewAdminSessionMiddleware(repo repository.SessionR) *AdminSession {
	return &AdminSession{
		sessionRepo: repo,
	}
}

func (a *AdminSession) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookies := c.Cookies()

		for _, cookie := range cookies {
			if cookie.Name == "admin_session" {
				if _, err := a.sessionRepo.GetSession(cookie.Value); err != nil {
					return errorHandler.ErrInvalidSession
				} else {
					return next(c)
				}
			}
		}

		return errorHandler.ErrInvalidSession
	}
}
