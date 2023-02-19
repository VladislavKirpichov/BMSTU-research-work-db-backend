package main

import (
	"fmt"
	"log"

	"github.com/v.kirpichov/admin/configs"
	network "github.com/v.kirpichov/admin/internal/network"
	"github.com/v.kirpichov/admin/internal/network/handlers"
	"github.com/v.kirpichov/admin/internal/network/middleware"
	repository2 "github.com/v.kirpichov/admin/internal/repository"
	"github.com/v.kirpichov/admin/internal/usecase"
	"github.com/v.kirpichov/admin/pkg/repository"
)

func main() {
	var config configs.Config
	err := configs.InitConfig(&config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresRepository(&config.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	redisUsersSessions, err := repository.NewRedisRepository(&config.RedisConfig)
	if err != nil {
		log.Fatal(err)
	}

	redisAdminSessions, err := repository.NewRedisRepository(&config.AdminRedisConfig)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository2.NewRepository(db, redisUsersSessions, redisAdminSessions)
	useCases := usecase.NewUsecases(repo, &config)
	handl := handlers.NewHandlers(useCases, &config)

	middlewares := middleware.New(repo)

	err = repository2.MustInitAdmins(db)
	if err != nil {
		log.Fatal(fmt.Errorf("Init admin error: %w", err))
	}

	e := network.InitRoutes(handl, middlewares)

	e.Logger.Fatal(e.Start(":8080"))
}
