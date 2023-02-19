package echoLogger

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func RequestLogger() echo.MiddlewareFunc {
	log := &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
	}

	log.SetFormatter(&prefixed.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Status >= 400 {
				log.WithFields(logrus.Fields{
					"URI":    v.URI,
					"Status": v.Status,
				}).Error("request")
			} else {
				log.WithFields(logrus.Fields{
					"Status": v.Status,
					"URI":    v.URI,
				}).Info("request")
			}

			return nil
		},
	})
}
