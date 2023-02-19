package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/internal/repository"
)

type Middleware struct {
	Session      SessionM
	AdminSession SessionM
}

func New(repo *repository.Repository) *Middleware {
	return &Middleware{
		Session:      NewSessionMiddleware(repo.Sessions),
		AdminSession: NewAdminSessionMiddleware(repo.AdminSessions),
	}
}

type SessionM interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}
