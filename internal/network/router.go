package network

import (
	"github.com/v.kirpichov/admin/pkg/errorHandler"

	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/internal/network/handlers"
	"github.com/v.kirpichov/admin/internal/network/middleware"
	"github.com/v.kirpichov/admin/pkg/echoLogger"
)

func InitRoutes(handlers *handlers.Handlers, middleware *middleware.Middleware) *echo.Echo {
	router := echo.New()

	router.Use(echoLogger.RequestLogger())
	router.HTTPErrorHandler = errorHandler.New().Handler

	apiGroup := router.Group("/api")
	apiGroup.GET("/users", handlers.UserHandler.GetAllUsers, middleware.Session.Auth)

	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/signin", handlers.UserHandler.SignIn)
	authGroup.POST("/signup", handlers.UserHandler.SignUp)

	adminGroup := apiGroup.Group("/admin")
	adminGroup.POST("/signin", handlers.AdminHandler.SignIn)
	adminGroup.GET("/users", handlers.UserHandler.GetAllUsers, middleware.AdminSession.Auth)

	servicesGroup := adminGroup.Group("/services")
	servicesGroup.GET("/:id", handlers.ServicesHandler.GetService, middleware.AdminSession.Auth)
	servicesGroup.GET("/", handlers.ServicesHandler.GetServices, middleware.AdminSession.Auth)
	servicesGroup.POST("/", handlers.ServicesHandler.CreateService, middleware.AdminSession.Auth)
	servicesGroup.PUT("/", handlers.ServicesHandler.UpdateService, middleware.AdminSession.Auth)
	servicesGroup.DELETE("/:id", handlers.ServicesHandler.DeleteService, middleware.AdminSession.Auth)

	employersGroup := adminGroup.Group("/employers")
	employersGroup.GET("/:id", handlers.EmployersHandler.GetEmployer, middleware.AdminSession.Auth)
	employersGroup.GET("/", handlers.EmployersHandler.GetEmployers, middleware.AdminSession.Auth)
	employersGroup.POST("/", handlers.EmployersHandler.CreateEmployer, middleware.AdminSession.Auth)
	employersGroup.PUT("/", handlers.EmployersHandler.UpdateEmployer, middleware.AdminSession.Auth)
	employersGroup.DELETE("/:id", handlers.EmployersHandler.DeleteEmployer, middleware.AdminSession.Auth)

	return router
}
