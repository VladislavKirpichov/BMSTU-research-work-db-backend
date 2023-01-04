package main

import (
	"log"

	"github.com/v.kirpichov/admin/configs"
	network "github.com/v.kirpichov/admin/internal/network"
	"github.com/v.kirpichov/admin/internal/network/handlers"
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

	redis, err := repository.NewRedisRepository(&config.RedisConfig)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository2.NewRepository(db, redis)
	useCases := usecase.NewUsecases(repo, &config)
	handl := handlers.NewHandlers(useCases, &config)

	e := network.InitRoutes(handl)

	e.Logger.Fatal(e.Start(":8080"))
}
