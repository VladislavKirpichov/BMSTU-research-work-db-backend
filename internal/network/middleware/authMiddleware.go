package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/siruspen/logrus"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logrus.Debug("hello world")

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
