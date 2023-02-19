package repository

import (
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/v.kirpichov/admin/configs"
)

func NewRedisRepository(cfg *configs.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     strings.Join([]string{cfg.Addr, strconv.Itoa(cfg.Port)}, ":"),
		Password: cfg.Password,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
