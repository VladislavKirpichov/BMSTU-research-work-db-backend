package network

import (
	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/internal/network/handlers"
)

func InitRoutes(handlers *handlers.Handlers) *echo.Echo {
	router := echo.New()

	router.GET("/signin", handlers.UserHandler.SignIn)
	router.GET("/users", handlers.UserHandler.GetAllUsers)
	router.POST("/signup", handlers.UserHandler.SignUp)

	return router
}
