package main

import (
	"log"

	"github.com/v.kirpichov/admin/configs"
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

	if db != nil {
		log.Default()
	}
}
