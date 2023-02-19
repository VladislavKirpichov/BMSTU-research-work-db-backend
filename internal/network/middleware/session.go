package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/internal/repository"
	"github.com/v.kirpichov/admin/pkg/errorHandler"
)

type Session struct {
	sessionsRepo repository.SessionR
}

func NewSessionMiddleware(repo repository.SessionR) *Session {
	return &Session{
		sessionsRepo: repo,
	}
}

func (s *Session) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookies := c.Cookies()

		for _, cookie := range cookies {
			if cookie.Name == "session" {
				fmt.Println(cookie.Value)
				if _, err := s.sessionsRepo.GetSession(cookie.Value); err != nil {
					return errorHandler.ErrInvalidSession
				} else {
					return next(c)
				}
			}
		}

		return errorHandler.ErrInvalidSession
	}
}
